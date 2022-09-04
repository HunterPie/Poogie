package middlewares

import (
	"net/http"

	"github.com/Haato3o/poogie/core/auth"
	"github.com/gin-gonic/gin"
)

const (
	TOKEN              = "X-Token"
	TRANSFORMED_HEADER = "X-Transformed-User-Id"
)

type UserTransformMiddleware struct {
	service auth.IAuthService
}

type UnauthorizedResponse struct {
	Error string `json:"error"`
}

func NewUserTransformMiddleware(service auth.IAuthService) *UserTransformMiddleware {
	return &UserTransformMiddleware{service}
}

func (m *UserTransformMiddleware) TokenToUserIdTransform(ctx *gin.Context) {
	ctx.Request.Header.Del(TRANSFORMED_HEADER)

	token := ctx.GetHeader(TOKEN)

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedResponse{
			Error: "You must log in first",
		})
		return
	}

	isValid := m.service.IsValid(token)

	if !isValid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedResponse{
			Error: "Invalid session token",
		})
		return
	}

	user, err := m.service.Parse(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedResponse{
			Error: "Invalid session token",
		})
		return
	}

	ctx.Request.Header.Add(TRANSFORMED_HEADER, user.UserId)

	ctx.Next()
}
