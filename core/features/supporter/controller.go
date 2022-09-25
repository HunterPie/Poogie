package supporter

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/Haato3o/poogie/core/features/common"
	"github.com/Haato3o/poogie/core/persistence/supporter"
	"github.com/Haato3o/poogie/core/services"
	"github.com/Haato3o/poogie/core/tracing"
	"github.com/Haato3o/poogie/core/utils"
	"github.com/Haato3o/poogie/pkg/http"
	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidWebhookCall = errors.New("invalid webhook call")
)

type SupporterController struct {
	service        *SupporterService
	patreonService *services.PatreonService
}

func (c *SupporterController) HandleSupporterWebhook(ctx *gin.Context) {
	txn := tracing.FromContext(ctx)
	body, err := io.ReadAll(ctx.Request.Body)
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	signature := ctx.Request.Header.Get("X-Patreon-Signature")
	event := ctx.Request.Header.Get("X-Patreon-Event")

	txn.AddProperty("event_type", event)

	if err != nil || !c.patreonService.IsWebhookValid(signature, body) {
		txn.AddProperty("error_type", "INVALID_PAYLOAD_SIGNATURE")
		http.BadRequest(ctx, common.ErrInvalidWebhook)
		return
	}

	webhook, err := c.patreonService.GetSupporterWebhook(ctx)

	if err != nil {
		txn.AddProperty("error_type", "INVALID_WEBHOOK_FORMAT")
		http.BadRequest(ctx, common.ErrInvalidWebhook)
		return
	}

	model, err := c.handleSupporterWebhookByType(ctx, event, webhook)

	if err != nil {
		txn.AddProperty("error_type", "INVALID_WEBHOOK_EVENT")
		http.BadRequest(ctx, common.ErrInvalidWebhook)
		return
	}

	http.Ok(ctx, SupporterResponse{
		Email:    model.Email,
		Token:    model.Token,
		IsActive: model.IsActive,
	})
}

func (c *SupporterController) HandleVerifySupporter(ctx *gin.Context) {
	var supporterHeader SupporterHeaderModel

	if !utils.DeserializeHeaders(ctx, &supporterHeader, func(header *SupporterHeaderModel) bool {
		return header.ClientId != ""
	}) {
		http.BadRequest(ctx, common.ErrGeneric)
		return
	}

	if supporterHeader.SupporterToken == "" {
		http.Ok(ctx, SupporterValidResponse{false})
		return
	}

	exists := c.service.ExistsSupporterByToken(ctx, supporterHeader.SupporterToken)

	http.Ok(ctx, SupporterValidResponse{exists})
}

func (c *SupporterController) handleSupporterWebhookByType(ctx context.Context, typ string, webhook services.PatreonWebhookModel) (supporter.SupporterModel, error) {
	txn := tracing.FromContext(ctx)
	email := webhook.Data.Attributes.Email
	txn.AddProperty("email", email)

	switch typ {
	case "members:pledge:create":
		return c.service.CreateNewSupporter(ctx, email), nil
	case "members:pledge:delete":
		return c.service.RevokeExistingSupporter(ctx, email), nil
	default:
		return supporter.SupporterModel{}, ErrInvalidWebhookCall
	}
}
