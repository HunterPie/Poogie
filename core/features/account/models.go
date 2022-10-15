package account

import (
	"time"

	"github.com/Haato3o/poogie/core/persistence/account"
)

type AccountCreationRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AvatarUpdateRequest struct {
	AvatarUrl string `json:"avatar_url"`
}

type AccountBadgeResponse struct {
	Id        string
	CreatedAt time.Time
}

type MyAccountResponse struct {
	Username    string                 `json:"username"`
	Email       string                 `json:"email"`
	AvatarUrl   string                 `json:"avatar_url"`
	Experience  int64                  `json:"experience"`
	Rating      int64                  `json:"rating"`
	Badges      []AccountBadgeResponse `json:"badges"`
	IsSupporter bool                   `json:"is_supporter"`
}

type UserAccountResponse struct {
	Username    string                 `json:"username"`
	AvatarUrl   string                 `json:"avatar_url"`
	Experience  int64                  `json:"experience"`
	Rating      int64                  `json:"rating"`
	Badges      []AccountBadgeResponse `json:"badges"`
	IsSupporter bool                   `json:"is_supporter"`
}

type AccountActivateResponse struct {
	Message string `json:"message"`
}

func toUserAccountResponse(account account.AccountModel) UserAccountResponse {
	return UserAccountResponse{
		Username:    account.Username,
		AvatarUrl:   account.AvatarUri,
		Experience:  account.Experience,
		Rating:      account.Rating,
		Badges:      toBadgesResponse(account.Badges),
		IsSupporter: account.IsSupporter,
	}
}

func toMyAccountResponse(account account.AccountModel) MyAccountResponse {
	return MyAccountResponse{
		Username:    account.Username,
		Email:       account.Email,
		AvatarUrl:   account.AvatarUri,
		Experience:  account.Experience,
		Rating:      account.Rating,
		Badges:      toBadgesResponse(account.Badges),
		IsSupporter: account.IsSupporter,
	}
}

func toBadgesResponse(badges []account.AccountBadgesModel) []AccountBadgeResponse {
	badgesResponse := make([]AccountBadgeResponse, len(badges))

	for _, badge := range badges {
		badgesResponse = append(badgesResponse, AccountBadgeResponse{
			Id:        badge.Id,
			CreatedAt: badge.CreatedAt,
		})
	}

	return badgesResponse
}

type PasswordResetRequest struct {
	Email string `json:"email"`
}

type PasswordResetResponse struct {
	Success bool `json:"success"`
}

type ChangePasswordRequest struct {
	Email       string `json:"email"`
	Code        string `json:"code"`
	NewPassword string `json:"new_password"`
}

type ChangePasswordResponse struct {
	Success bool `json:"success"`
}
