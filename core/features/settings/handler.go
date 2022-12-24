package settings

import (
	"github.com/Haato3o/poogie/core/middlewares"
	"github.com/Haato3o/poogie/pkg/server"
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h Handler) Load(router *gin.RouterGroup, server *server.Server) error {
	service := Service{
		repository: server.Database.GetClientSettingsRepository(),
	}
	controller := Controller{
		s: &service,
	}

	clientRouter := router.Group("/account/client")

	clientRouter.Use(middlewares.BlockUnauthenticatedRequest)

	clientRouter.GET("/settings", controller.GetClientSettingsHandler)
	clientRouter.PATCH("/settings", controller.UpdateClientSettingsHandler)

	return nil
}

func (h Handler) GetVersion() int {
	return server.V1
}

func (h Handler) GetName() string {
	return "SettingsHandler"
}

func New() server.IRegisterableService {
	return &Handler{}
}
