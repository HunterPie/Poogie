package account

import (
	"github.com/Haato3o/poogie/core/utils"
	"github.com/Haato3o/poogie/pkg/http"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	service *AccountService
}

func (c *AccountController) CreateNewAccountHandler(ctx *gin.Context) {
	var request AccountCreationRequest
	ok := utils.DeserializeBody(ctx, &request)

	if !ok {
		http.BadRequest(ctx)
		return
	}

	clientId := utils.ExtractClientId(ctx)

	account, err := c.service.CreateNewAccount(ctx, request, clientId)

	if err != nil {
		http.InternalServerError(ctx)
		return
	}

	http.Ok(ctx, toAccountResponse(account))
}

func (c *AccountController) GetUserHandler(ctx *gin.Context) {
	userId := ctx.Param("userId")

	account, err := c.service.repository.GetById(ctx, userId)

	if err != nil {
		http.ElementNotFound(ctx)
		return
	}

	http.Ok(ctx, toUserAccountResponse(account))
}

func (c *AccountController) GetMyUserHandler(ctx *gin.Context) {
	userId := utils.ExtractUserId(ctx)

	account, _ := c.service.repository.GetById(ctx, userId)

	http.Ok(ctx, toAccountResponse(account))
}
