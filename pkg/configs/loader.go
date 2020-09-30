package configs

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"ped-consul-watch/pkg/api"
	"ped-consul-watch/pkg/provider"
	"strings"
)

const (
	orgsPath   = "configs/orgs"
)

type loader struct {
	ce provider.ConsulExecutor
}

func NewConsulLoader(c provider.ConsulExecutor) Loader {
	return loader{
		ce: c,
	}
}

func (l loader) LoadOrganizations() (*api.Configs, error) {
	orgs, err := l.ce.KVList(orgsPath)
	if err != nil {
		log.Err(err).Msgf("Fail load orgs %s", orgsPath)
		return nil, err
	}
	orgsConfigs := make(map[string]api.OrgConfig, len(orgs))
	for k, v := range orgs {
		var config api.OrgConfig
		err := json.Unmarshal(v, &config)
		if err != nil {
			log.Fatal().Err(err).Msgf("Fail unmarshal config consul %s", k)
		}
		org := strings.ReplaceAll(k, "configs/orgs/", "")
		orgsConfigs[org] = config
	}
	return &api.Configs{
		OrgConfigs: orgsConfigs,
	}, nil
}
