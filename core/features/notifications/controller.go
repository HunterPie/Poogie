package notifications

import (
	"github.com/Haato3o/poogie/pkg/http"
	"github.com/gin-gonic/gin"
)

type NotificationsController struct {
	*NotificationsService
}

func (c *NotificationsController) GetAllNotificationsHandler(ctx *gin.Context) {
	notifications := c.GetAllNotifications(ctx)
	http.Ok(ctx, notifications)
}
