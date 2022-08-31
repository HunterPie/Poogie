package supporter

import (
	"github.com/Haato3o/poogie/core/services"
	"github.com/Haato3o/poogie/core/utils"
	"github.com/Haato3o/poogie/pkg/http"
	"github.com/gin-gonic/gin"
)

type SupporterController struct {
	service        *SupporterService
	patreonService *services.PatreonService
}

func (c *SupporterController) HandleCreateSupporterWebhook(ctx *gin.Context) {
	webhook, err := c.patreonService.GetSupporterWebhook(ctx)

	if err != nil {
		http.BadRequest(ctx)
		return
	}

	model := c.service.CreateNewSupporter(ctx, webhook.Data.Email)

	http.Ok(ctx, SupporterResponse{
		Email:    model.Email,
		Token:    model.Token,
		IsActive: model.IsActive,
	})
}

func (c *SupporterController) VerifySupporter(ctx *gin.Context) {
	var supporterHeader SupporterHeaderModel

	if !utils.DeserializeHeaders(ctx, &supporterHeader, func(header *SupporterHeaderModel) bool {
		return header.SupporterToken != ""
	}) {
		http.BadRequest(ctx)
		return
	}

	exists := c.service.ExistsSupporterByToken(ctx, supporterHeader.SupporterToken)

	http.Ok(ctx, SupporterValidResponse{exists})
}
