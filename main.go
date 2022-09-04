package main

import (
	"fmt"

	"github.com/Haato3o/poogie/core/config"
	"github.com/Haato3o/poogie/core/features/account"
	"github.com/Haato3o/poogie/core/features/health"
	"github.com/Haato3o/poogie/core/features/report"
	"github.com/Haato3o/poogie/core/features/session"
	"github.com/Haato3o/poogie/core/features/supporter"
	"github.com/Haato3o/poogie/core/features/version"
	"github.com/Haato3o/poogie/pkg/log"
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
		report.New(),
		session.New(),
		account.New(),
	}
}

func groupByVersions(services []server.IRegisterableService) map[int][]server.IRegisterableService {
	m := make(map[int][]server.IRegisterableService, 0)

	for _, svc := range services {
		version := svc.GetVersion()

		_, ok := m[version]

		if !ok {
			m[version] = make([]server.IRegisterableService, 0)
		}

		m[version] = append(m[version], svc)
	}

	return m
}

func main() {
	godotenv.Load()
	_ = envconfig.Process("POOGIE", &apiConfig)

	instance, err := server.New(&apiConfig)

	if err != nil {
		panic("failed to instantiate new server")
	}

	groupedServices := groupByVersions(getServices())
	err = instance.Load(groupedServices)

	if err != nil {
		panic(err)
	}

	log.Info(fmt.Sprintf("Starting up server at %s", apiConfig.HttpAddress))

	instance.Start()
}
