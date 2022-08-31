package main

import (
	"fmt"

	"github.com/Haato3o/poogie/core/config"
	"github.com/Haato3o/poogie/core/features/health"
	"github.com/Haato3o/poogie/pkg/server"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var apiConfig config.ApiConfiguration

func getServices() []server.IRegisterableService {
	return []server.IRegisterableService{
		health.New(),
	}
}

func groupByVersions(services []server.IRegisterableService) map[int][]server.IRegisterableService {
	m := make(map[int][]server.IRegisterableService)

	for _, svc := range services {
		version := svc.GetVersion()

		_, ok := m[version]

		if !ok {
			m[version] = make([]server.IRegisterableService, 1)
		}

		m[version] = append(m[version], svc)
	}

	return m
}

func loadAllServices()

func main() {
	godotenv.Load()
	_ = envconfig.Process("POOGIE", &apiConfig)

	// TODO: Add metrics

	instance, err := server.New(&apiConfig)

	if err != nil {
		panic("failed to instantiate new server")
	}

	for version, services := range groupByVersions(getServices()) {
		var group string = ""
		if version > server.NO_VERSION {
			group = fmt.Sprintf("v%d", version)
		}

		router := instance.HttpServer.Router.Group(group)

		for _, service := range services {
			err := service.Load(router, instance)

			if err != nil {
				panic("failed to load handler")
			}
		}
	}

}
