package supporter

type SupporterHeaderModel struct {
	ClientId       string `header:"X-Client-Id"`
	SupporterToken string `header:"X-Supporter-Token"`
}

type SupporterResponse struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	IsActive bool   `json:"is_active"`
}

type SupporterValidResponse struct {
	IsValid bool `json:"is_valid"`
}
