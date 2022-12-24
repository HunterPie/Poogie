package settings

import (
	"github.com/Haato3o/poogie/core/features/common"
	"github.com/Haato3o/poogie/core/utils"
	"github.com/Haato3o/poogie/pkg/http"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	s *Service
}

func (c *Controller) GetClientSettingsHandler(ctx *gin.Context) {
	clientId := utils.ExtractClientId(ctx)

	config, err := c.s.GetSettings(ctx, clientId)

	if err != nil {
		handleException(ctx, err)
		return
	}

	http.Ok(ctx, toClientSettingsResponse(config))
}

func (c *Controller) UpdateClientSettingsHandler(ctx *gin.Context) {
	clientId := utils.ExtractUserId(ctx)

	var request UpdateClientSettingsRequest
	ok, _ := utils.DeserializeBody(ctx, &request)

	if !ok {
		http.BadRequest(ctx, common.ErrInvalidPayload)
		return
	}

	config, err := c.s.UpdateSettings(ctx, clientId, request.Configuration)

	if err != nil {
		handleException(ctx, err)
		return
	}

	http.Ok(ctx, toClientSettingsResponse(config))
}

func handleException(ctx *gin.Context, err error) {
	if err == ErrInvalidUser {
		http.ElementNotFound(ctx)
	} else if err == ErrInvalidSettings {
		http.BadRequest(ctx, common.ErrInvalidPayload)
	} else if err != nil {
		http.InternalServerError(ctx)
	}
}
