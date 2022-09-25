package supporter

type SupporterModel struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	IsActive bool   `json:"is_active"`
}
