package settings

import "github.com/Haato3o/poogie/core/persistence/settings"

type ClientSettingsResponse struct {
	Configuration string `json:"configuration"`
}

func toClientSettingsResponse(m settings.ClientSettingModel) ClientSettingsResponse {
	return ClientSettingsResponse{
		Configuration: m.EncodedSettings,
	}
}

type UpdateClientSettingsRequest struct {
	Configuration string `json:"configuration" binding:"required"`
}
