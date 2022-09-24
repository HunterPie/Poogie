package save

import "time"

type GameType string

const (
	GameTypeMHW = "MHW"
	GameTypeMHR = "MHR"
)

type BackupModel struct {
	Id        string    `json:"id"`
	Type      GameType  `json:"type"`
	CreatedAt time.Time `json:"create_at"`
}

type SaveModel struct {
	UserId  string        `json:"user_id"`
	Backups []BackupModel `json:"backups"`
}
