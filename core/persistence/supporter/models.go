package supporter

type SupporterModel struct {
	Email    string `json:"email" bson:"email"`
	Token    string `json:"token" bson:"token"`
	IsActive bool   `json:"is_active" bson:"is_active"`
}
