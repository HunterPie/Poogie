package main

import (
	"github.com/Haato3o/poogie/core/config"
	"github.com/Haato3o/poogie/core/features/health"
	"github.com/Haato3o/poogie/core/features/supporter"
	"github.com/Haato3o/poogie/core/features/version"
	"github.com/Haato3o/poogie/pkg/server"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var apiConfig config.ApiConfiguration

func getServices() []server.IRegisterableService {
	return []server.IRegisterableService{
		health.New(),
		version.New(),
		supporter.New(),
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

func main() {
	godotenv.Load()
	_ = envconfig.Process("POOGIE", &apiConfig)

	// TODO: Add metrics

	instance, err := server.New(&apiConfig)

	if err != nil {
		panic("failed to instantiate new server")
	}

	groupedServices := groupByVersions(getServices())
	instance.Load(groupedServices)

	instance.Start()
}
