package localization

type GetLocalizationsResponse struct {
	Localizations map[string]string `json:"localizations"`
}
