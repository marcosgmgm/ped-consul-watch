package watch

import (
	"encoding/base64"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"net/http"
	"ped-consul-watch/pkg/api"
	"ped-consul-watch/pkg/dhttp"
	"strings"
)

type Handler struct {
	configs *api.Configs
}

func NewHttpHandler(c *api.Configs) dhttp.HttpHandler{
	return Handler{
		configs: c,
	}
}

type request struct {
	Key string `json:"Key"`
	CreateIndex int `json:"CreateIndex"`
	ModifyIndex int `json:"ModifyIndex"`
	LockIndex int `json:"LockIndex"`
	Flags int `json:"Flags"`
	Value string `json:"Value"`
	Session string `json:"Session"`
}

func (h Handler) Handler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		log.Info().Msg("Receive request watch")
		ch := dhttp.ConfigsFromRequest(r)
		log.Info().Msgf("x-configs : {%s}", ch)
		var req []request
		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Error().Err(err).Msg("error read body")
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}
		for _, re := range req {
			org := strings.ReplaceAll(re.Key, "configs/orgs/", "")
			log.Info().Msgf("receive watch to org : {%s}", org)
			vb, err := base64.StdEncoding.DecodeString(re.Value)
			if err != nil {
				log.Error().Err(err).Msg("error decode value")
				http.Error(w, "can't decode value", http.StatusBadRequest)
				return
			}
			var orgConfig api.OrgConfig
			json.Unmarshal(vb, &orgConfig)
			log.Info().Msgf("receive watch with org config : {%v}", orgConfig)

			log.Info().Msgf("Updating config to org [%s]", org)
			h.configs.OrgConfigs[org] = orgConfig
			log.Info().Msgf("Config updated to org [%s]", org)
		}
	}
}
