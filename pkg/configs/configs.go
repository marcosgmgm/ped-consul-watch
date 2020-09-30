package configs

import "ped-consul-watch/pkg/api"

type Loader interface {
	LoadOrganizations() (*api.Configs, error)
}
