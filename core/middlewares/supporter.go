package middlewares

import (
	"fmt"
	"github.com/Haato3o/poogie/core/persistence/account"
	"github.com/Haato3o/poogie/core/persistence/database"
	"github.com/Haato3o/poogie/core/persistence/supporter"
	"github.com/Haato3o/poogie/core/utils"
	"github.com/gin-gonic/gin"
)

const (
	TransformedSupporterHeader = "X-Transformed-Is-Supporter"
)

type SupporterTransformMiddleware struct {
	accountRepository   account.IAccountRepository
	supporterRepository supporter.ISupporterRepository
}

func (m *SupporterTransformMiddleware) TransformSupporterToken(ctx *gin.Context) {
	ctx.Request.Header.Del(TransformedSupporterHeader)

	token := utils.ExtractSupporterToken(ctx)
	userId := utils.ExtractUserId(ctx)

	isSupporter := false

	if token != "" {
		isSupporter = isSupporter || m.supporterRepository.ExistsToken(ctx, token)
	}

	if userId != "" && !isSupporter {
		accountInfo, _ := m.accountRepository.GetById(ctx, userId)
		isSupporter = isSupporter || accountInfo.IsSupporter
	}

	ctx.Request.Header.Add(TransformedSupporterHeader, fmt.Sprintf("%t", isSupporter))

	ctx.Next()
}

func NewSupporterTransformMiddleware(database database.IDatabase) *SupporterTransformMiddleware {
	return &SupporterTransformMiddleware{
		accountRepository:   database.GetAccountRepository(),
		supporterRepository: database.GetSupporterRepository(),
	}
}
