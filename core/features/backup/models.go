package backup

import (
	"time"

	"github.com/Haato3o/poogie/core/persistence/backups"
)

const (
	GAME_MONSTER_HUNTER_RISE  = "Monster Hunter Rise: Sunbreak"
	GAME_MONSTER_HUNTER_WORLD = "Monster Hunter: World"

	GAME_MONSTER_HUNTER_RISE_ICON  = "https://cdn.hunterpie.com/Static/monster-hunter-rise-icon.png"
	GAME_MONSTER_HUNTER_WORLD_ICON = "https://cdn.hunterpie.com/Static/monster-hunter-world-icon.png"
)

type UserBackupDetailsResponse struct {
	Count    int              `json:"count"`
	MaxCount int              `json:"max_count"`
	Backups  []BackupResponse `json:"backups"`
}

type BackupResponse struct {
	Id         string    `json:"id"`
	Size       int64     `json:"size"`
	Game       string    `json:"game_name"`
	Icon       string    `json:"game_icon"`
	UploadedAt time.Time `json:"uploaded_at"`
}

func ToBackupResponses(models []backups.BackupUploadModel) []BackupResponse {
	responses := make([]BackupResponse, 0)

	for _, model := range models {
		responses = append(responses, ToBackupResponse(model))
	}

	return responses
}

func ToBackupResponse(model backups.BackupUploadModel) BackupResponse {
	return BackupResponse{
		Id:         model.Id,
		Size:       model.Size,
		Game:       toGameName(model.Game),
		Icon:       toGameIcon(model.Game),
		UploadedAt: model.UploadedAt,
	}
}

func toGameName(gameType backups.GameType) string {
	switch gameType {
	case backups.MHR:
		return GAME_MONSTER_HUNTER_RISE
	case backups.MHW:
		return GAME_MONSTER_HUNTER_WORLD
	default:
		return ""
	}
}

func toGameIcon(gameType backups.GameType) string {
	switch gameType {
	case backups.MHR:
		return GAME_MONSTER_HUNTER_RISE_ICON
	case backups.MHW:
		return GAME_MONSTER_HUNTER_WORLD_ICON
	default:
		return ""
	}
}
