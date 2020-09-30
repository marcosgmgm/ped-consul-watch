package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"ped-consul-watch/pkg/api"
	"ped-consul-watch/pkg/configs"
	"ped-consul-watch/pkg/provider"
	"ped-consul-watch/pkg/watch"
	"time"
)

func main() {
	// zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Debug().Msg("start app...")

	//Consul provider
	consulExecutor, err := provider.NewConsulExecutor()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create consul executor")
	}

	//Config Loader
	configLoader := configs.NewConsulLoader(consulExecutor)

	//Consul configuration key /configs/orgs/*
	configs, err := configLoader.LoadOrganizations()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load consul configuration")
	}

	//print config
	go printConfigs(configs)

	//Handler
	wh := watch.NewHttpHandler(configs)

	//Routers
	router := httprouter.New()
	router.POST("/watch", wh.Handler())
	log.Fatal().Err(http.ListenAndServe(":3000", router))
}

func printConfigs(config *api.Configs)  {
	for {
		for k, v := range config.OrgConfigs {
			log.Info().Msgf("print configs org: {%s} - OrgConfig: name: {%s}, database :{%s}", k, v.Name, v.Database)
		}
		time.Sleep(time.Second * 15)
	}
}
