package settings

import "context"

type IClientSettingsRepository interface {
	UpdateBy(ctx context.Context, userId string, encodedSettings string) (ClientSettingModel, error)
	DeleteBy(ctx context.Context, userId string) bool
	FindBy(ctx context.Context, userId string) (ClientSettingModel, error)
}
