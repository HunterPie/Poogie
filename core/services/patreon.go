package services

import (
	"errors"

	"github.com/Haato3o/poogie/core/utils"
	"github.com/gin-gonic/gin"
)

type PatreonAttributeModel struct {
	Email string `json:"email" binding:"required"`
}

type PatreonDataWebhookModel struct {
	Attributes PatreonAttributeModel `json:"attributes" binding:"required"`
	Id         string                `json:"id"`
	Type       string                `json:"type"`
}

type PatreonWebhookModel struct {
	Data PatreonAttributeModel `json:"data" binding:"required"`
}

type PatreonService struct{}

func (s *PatreonService) GetSupporterWebhook(ctx *gin.Context) (PatreonWebhookModel, error) {
	var body PatreonWebhookModel

	valid := utils.DeserializeBody(ctx, &body)

	if !valid {
		return body, errors.New("wrong webhook payload format")
	}

	return body, nil
}
