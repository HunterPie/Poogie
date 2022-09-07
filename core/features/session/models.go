package session

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}

type CreateSessionResponse struct {
	Token string `json:"token"`
}
