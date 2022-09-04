package auth

type AuthPayload struct {
	UserId string `json:"user_id"`
}

type IAuthService interface {
	Create(userId string) string
	Refresh(token string) string
	IsValid(token string) bool
	Parse(token string) (AuthPayload, error)
}
