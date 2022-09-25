package middlewares

import (
	"net/http"

	"github.com/Haato3o/poogie/core/auth"
	"github.com/Haato3o/poogie/core/crypto"
	"github.com/Haato3o/poogie/core/features/common"
	"github.com/Haato3o/poogie/core/persistence/account"
	"github.com/gin-gonic/gin"
)

const (
	TOKEN              = "X-Token"
	TRANSFORMED_HEADER = "X-Transformed-User-Id"
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
	ctx.Request.Header.Del(TRANSFORMED_HEADER)

	token := ctx.GetHeader(TOKEN)

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedResponse{
			Error: "You must log in first",
			Code:  common.ErrInvalidSessionToken,
		})
		return
	}

	isValid := m.service.IsValid(token)

	if !isValid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedResponse{
			Error: "Invalid session token",
			Code:  common.ErrInvalidSessionToken,
		})
		return
	}

	user, err := m.service.Parse(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedResponse{
			Error: "Invalid session token",
			Code:  common.ErrInvalidSessionToken,
		})
		return
	}

	hashedToken := m.hashService.Hash(token)
	IsSessionValid := m.repository.IsSessionValid(ctx, hashedToken)

	if !IsSessionValid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedResponse{
			Error: "Invalid session token",
			Code:  common.ErrInvalidSessionToken,
		})
		return
	}

	ctx.Request.Header.Add(TRANSFORMED_HEADER, user.UserId)

	ctx.Next()
}
