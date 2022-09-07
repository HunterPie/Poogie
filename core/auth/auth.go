package auth

type AuthPayload struct {
	UserId string `json:"user_id"`
}

type IAuthService interface {
	Create(userId string) (string, error)
	IsValid(token string) bool
	Parse(token string) (AuthPayload, error)
}
