package supporter

type SupporterModel struct {
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	IsActive bool   `json:"is_active"`
}
