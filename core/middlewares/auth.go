package middlewares

import (
	"github.com/Haato3o/poogie/core/utils"
	"net/http"

	"github.com/Haato3o/poogie/core/auth"
	"github.com/Haato3o/poogie/core/crypto"
	"github.com/Haato3o/poogie/core/features/common"
	"github.com/Haato3o/poogie/core/persistence/account"
	"github.com/gin-gonic/gin"
)

const (
	TOKEN             = "X-Token"
	TransformedHeader = "X-Transformed-User-Id"
)

type UserTransformMiddleware struct {
	service     auth.IAuthService
	repository  account.IAccountSessionRepository
	hashService crypto.IHashService
}

type UnauthorizedResponse struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

func NewUserTransformMiddleware(service auth.IAuthService, repository account.IAccountSessionRepository, hashService crypto.IHashService) *UserTransformMiddleware {
	return &UserTransformMiddleware{service, repository, hashService}
}

func (m *UserTransformMiddleware) TokenToUserIdTransform(ctx *gin.Context) {
	ctx.Request.Header.Del(TransformedHeader)

	token := ctx.GetHeader(TOKEN)
	isValid := m.service.IsValid(token)
	userId := ""

	if token != "" && isValid {
		payload, _ := m.service.Parse(token)

		hashedToken := m.hashService.Hash(token)
		IsSessionValid := m.repository.IsSessionValid(ctx, hashedToken)

		if IsSessionValid {
			userId = payload.UserId
		}
	}

	ctx.Request.Header.Add(TransformedHeader, userId)

	ctx.Next()
}

func BlockUnauthenticatedRequest(ctx *gin.Context) {
	userId := utils.ExtractUserId(ctx)

	if userId == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedResponse{
			Error: "Invalid session token",
			Code:  common.ErrInvalidSessionToken,
		})
		return
	}

	ctx.Next()
}
