package dhttp

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const ConfigsHeader = "x-configs"

type HttpHandler interface {
	Handler() httprouter.Handle
}

func ConfigsFromRequest(r *http.Request) string {
	return r.Header.Get(ConfigsHeader)
}