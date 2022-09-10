package notifications

import (
	"github.com/Haato3o/poogie/pkg/server"
	"github.com/gin-gonic/gin"
)

type NotificationsHandler struct{}

// GetName implements server.IRegisterableService
func (*NotificationsHandler) GetName() string {
	return "NotificationsHandler"
}

// GetVersion implements server.IRegisterableService
func (*NotificationsHandler) GetVersion() int {
	return server.V1
}

// Load implements server.IRegisterableService
func (*NotificationsHandler) Load(router *gin.RouterGroup, server *server.Server) error {
	service := NotificationsService{
		repository: server.Database.GetNotificationsRepository(),
	}
	controller := NotificationsController{
		NotificationsService: &service,
	}

	router.GET("/notifications", controller.GetAllNotificationsHandler)

	return nil
}

func New() server.IRegisterableService {
	return &NotificationsHandler{}
}
