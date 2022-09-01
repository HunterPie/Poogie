package services

import (
	"crypto/hmac"
	"crypto/md5"
	"errors"
	"fmt"

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

type PatreonService struct {
	secret string
}

func NewPatreonService(secret string) *PatreonService {
	return &PatreonService{secret}
}

func (s *PatreonService) GetSupporterWebhook(ctx *gin.Context) (PatreonWebhookModel, error) {
	var body PatreonWebhookModel

	valid := utils.DeserializeBody(ctx, &body)

	if !valid {
		return body, errors.New("wrong webhook payload format")
	}

	return body, nil
}

func (s *PatreonService) IsWebhookValid(signature string, body []byte) bool {
	mac := hmac.New(md5.New, []byte(s.secret))
	mac.Write(body)
	stringfied := fmt.Sprintf("%x", mac.Sum(nil))

	return signature == stringfied
}
