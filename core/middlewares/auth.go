package middlewares

import (
	"net/http"

	"github.com/Haato3o/poogie/core/persistence/account"
	"github.com/Haato3o/poogie/core/persistence/database"
	"github.com/gin-gonic/gin"
)

const (
	TOKEN              = "X-Token"
	TRANSFORMED_HEADER = "X-Transformed-User-Id"
)

type UserTransformMiddleware struct {
	repository account.IAccountSessionRepository
}

type UnauthorizedResponse struct {
	Error string `json:"error"`
}

func NewUserTransformMiddleware(db database.IDatabase) *UserTransformMiddleware {
	return &UserTransformMiddleware{
		repository: db.GetSessionRepository(),
	}
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

	userId, err := m.repository.GetUserIdBy(ctx, token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedResponse{
			Error: "Invalid session token",
		})
	}

	ctx.Request.Header.Add(TRANSFORMED_HEADER, userId)

	ctx.Next()
}
