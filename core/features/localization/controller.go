package localization

import (
	"github.com/Haato3o/poogie/pkg/http"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *Service
}

func (c *Controller) GetLocalizationsHandler(ctx *gin.Context) {
	localizations := c.service.ListAvailableLocalizations(ctx)

	http.Ok(ctx, GetLocalizationsResponse{
		Localizations: localizations,
	})
}

func NewController(service *Service) *Controller {
	return &Controller{service: service}
}
