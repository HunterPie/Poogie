package main

import (
	"fmt"
	"github.com/Haato3o/poogie/core/features/localization"
	"github.com/Haato3o/poogie/core/features/settings"

	"github.com/Haato3o/poogie/core/config"
	"github.com/Haato3o/poogie/core/features/account"
	"github.com/Haato3o/poogie/core/features/backup"
	"github.com/Haato3o/poogie/core/features/health"
	"github.com/Haato3o/poogie/core/features/notifications"
	"github.com/Haato3o/poogie/core/features/report"
	"github.com/Haato3o/poogie/core/features/session"
	"github.com/Haato3o/poogie/core/features/supporter"
	"github.com/Haato3o/poogie/core/features/version"
	"github.com/Haato3o/poogie/core/services"
	"github.com/Haato3o/poogie/metadata"
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
		notifications.New(),
		backup.New(),
		localization.New(),
		settings.New(),
	}
}

func groupByVersions(services []server.IRegisterableService) map[int][]server.IRegisterableService {
	m := make(map[int][]server.IRegisterableService, 0)

	for _, svc := range services {
		svcVersion := svc.GetVersion()

		_, ok := m[svcVersion]

		if !ok {
			m[svcVersion] = make([]server.IRegisterableService, 0)
		}

		m[svcVersion] = append(m[svcVersion], svc)
	}

	return m
}

func main() {
	_ = godotenv.Load()
	_ = envconfig.Process("POOGIE", &apiConfig)

	log.NewLogger(apiConfig.NewRelicLicenseKey)

	webhookService := services.NewDiscordWebhookService(apiConfig.DeployWebhook)

	instance, err := server.New(&apiConfig)

	if err != nil {
		log.Error("failed to instantiate server", err)
		panic("failed to instantiate new server")
	}

	groupedServices := groupByVersions(getServices())
	err = instance.Load(groupedServices)

	if err != nil {
		panic(err)
	}

	log.Info(fmt.Sprintf("Starting up server at %s. Version %s", apiConfig.HttpAddress, metadata.Version))

	webhookService.SendEmbed(config.DeployEmbed)

	instance.Start()

	webhookService.SendEmbed(config.DiedEmbed)
}
