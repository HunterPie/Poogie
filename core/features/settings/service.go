package settings

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/Haato3o/poogie/core/persistence/settings"
)

var (
	ErrInvalidUser     = errors.New("invalid user")
	ErrInvalidSettings = errors.New("invalid settings")
)

type Service struct {
	repository settings.IClientSettingsRepository
}

func (s *Service) UpdateSettings(ctx context.Context, userId string, newConfig string) (settings.ClientSettingModel, error) {
	buffer := make([]byte, base64.StdEncoding.DecodedLen(len(newConfig)))

	_, err := base64.StdEncoding.Decode(buffer, []byte(newConfig))

	if err != nil {
		return settings.ClientSettingModel{}, ErrInvalidSettings
	}

	buffer = bytes.Trim(buffer, "\x00")

	if !json.Valid(buffer) {
		return settings.ClientSettingModel{}, ErrInvalidSettings
	}

	config, err := s.repository.UpdateBy(ctx, userId, newConfig)

	if err != nil {
		return settings.ClientSettingModel{}, ErrInvalidUser
	}

	return config, nil
}

func (s *Service) GetSettings(ctx context.Context, userId string) (settings.ClientSettingModel, error) {
	config, err := s.repository.FindBy(ctx, userId)

	if err != nil {
		return config, ErrInvalidUser
	}

	return config, nil
}
