package session

import (
	"github.com/Haato3o/poogie/core/features/common"
	"github.com/Haato3o/poogie/core/tracing"
	"github.com/Haato3o/poogie/core/utils"
	"github.com/Haato3o/poogie/pkg/http"
	"github.com/gin-gonic/gin"
)

type SessionController struct {
	service *SessionService
}

func (c *SessionController) LoginHandler(ctx *gin.Context) {
	txn := tracing.FromContext(ctx)
	var request LoginRequest
	ok, _ := utils.DeserializeBody(ctx, &request)

	if !ok {
		http.BadRequest(ctx, common.ErrInvalidPayload)
		return
	}

	token, err := c.service.CreateSession(ctx, request)

	if err == ErrWrongCredentials {
		http.Unauthorized(ctx)
		return
	}

	if err == ErrUnverifiedAccount {
		http.UnauthorizedWithCustomError(ctx, common.ErrUnverifiedAccount, "Account is not verified yet, please check your inbox and spam folders for the verification email.")
		return
	}

	if err != nil {
		txn.AddProperty("error_message", err)
		http.InternalServerError(ctx)
		return
	}

	http.Ok(ctx, CreateSessionResponse{
		Token: token,
	})
}

func (c *SessionController) LogoutHandler(ctx *gin.Context) {
	token := ctx.GetHeader("X-Token")

	if token == "" {
		http.Unauthorized(ctx)
		return
	}

	c.service.RevokeSession(ctx, token)

	http.Ok(ctx, LogoutResponse{
		Message: "Logged out",
	})
}
