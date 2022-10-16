package report

import (
	"github.com/Haato3o/poogie/core/services"
	"github.com/Haato3o/poogie/pkg/server"
	"github.com/gin-gonic/gin"
)

type ReportHandler struct{}

// GetName implements server.IRegisterableService
func (*ReportHandler) GetName() string {
	return "ReportHandler"
}

// GetVersion implements server.IRegisterableService
func (*ReportHandler) GetVersion() int {
	return server.V1
}

// Load implements server.IRegisterableService
func (*ReportHandler) Load(router *gin.RouterGroup, server *server.Server) error {
	service := ReportService{
		webhookService: services.NewDiscordWebhookService(server.Config.DiscordCrashWebhook),
	}
	controller := ReportController{
		service: &service,
	}

	router.POST("/report/crash", controller.SendCrashReportHandler)

	return nil
}

func New() server.IRegisterableService {
	return &ReportHandler{}
}
