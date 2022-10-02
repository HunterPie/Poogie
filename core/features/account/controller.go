package account

import (
	"bytes"
	"io"

	"github.com/Haato3o/poogie/core/features/common"
	"github.com/Haato3o/poogie/core/utils"
	"github.com/Haato3o/poogie/pkg/http"
	"github.com/gin-gonic/gin"
)

const MAX_AVATAR_SIZE = 3500000

type AccountController struct {
	service *AccountService
}

func (c *AccountController) CreateNewAccountHandler(ctx *gin.Context) {
	var request AccountCreationRequest
	ok, handled := utils.DeserializeBody(ctx, &request, func(t *AccountCreationRequest) (bool, bool) {

		if !utils.ValidateEmail(t.Email) {
			http.BadRequest(ctx, common.ErrInvalidEmail)
			return false, true
		}

		if len(request.Username) < 3 {
			http.BadRequest(ctx, common.ErrInvalidUsername)
			return false, true
		}

		if len(request.Password) < 8 {
			http.BadRequest(ctx, common.ErrInvalidPassword)
			return false, true
		}

		return true, false
	})

	if handled {
		return
	}

	if !ok {
		http.BadRequest(ctx, common.ErrInvalidPayload)
		return
	}

	clientId := utils.ExtractClientId(ctx)

	account, err := c.service.CreateNewAccount(ctx, request, clientId)

	if err == ErrAccountWithEmailAlreadyExists || err == ErrUsernameTaken {
		http.Conflict(ctx, err.Error(), common.ErrUserAlreadyExists)
		return
	} else if err != nil {
		http.InternalServerError(ctx)
		return
	}

	http.Ok(ctx, toAccountResponse(account))
}

func (c *AccountController) VerifyAccount(ctx *gin.Context) {
	token := ctx.Param("token")

	_, err := c.service.VerifyAccount(ctx, token)

	if err == ErrAlreadyActivated {
		http.Ok(ctx, AccountActivateResponse{
			Message: "That account has already been verified",
		})
		return
	} else if err != nil {
		http.ElementNotFound(ctx)
		return
	}

	http.Ok(ctx, AccountActivateResponse{
		Message: "Account is now verified!",
	})
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

	response := toAccountResponse(account)

	response.Email, _ = c.service.cryptoService.Decrypt(response.Email)

	http.Ok(ctx, response)
}

func (c *AccountController) UploadAvatar(ctx *gin.Context) {
	userId := utils.ExtractUserId(ctx)

	file, headers, err := ctx.Request.FormFile("file")

	if err != nil {
		http.BadRequest(ctx, common.ErrInvalidImage)
		return
	}

	if headers.Size > MAX_AVATAR_SIZE {
		http.TooLarge(ctx, common.ErrAvatarSizeTooLarge)
		return
	}

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, file)

	if err != nil {
		http.InternalServerError(ctx)
		return
	}

	account, err := c.service.UpdateAvatar(ctx, userId, buffer.Bytes())

	if err != nil {
		http.BadRequest(ctx, common.ErrAvatarUploadFail)
		return
	}

	http.Ok(ctx, toUserAccountResponse(account))
}
