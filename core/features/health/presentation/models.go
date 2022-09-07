package presentation

type IsHealthyResponse struct {
	Healthy bool   `json:"is_alive"`
	Error   string `json:"error,omitempty"`
}
