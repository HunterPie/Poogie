package session

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateSessionResponse struct {
	Token string `json:"token"`
}
