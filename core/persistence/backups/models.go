package backups

import "time"

const (
	MHR GameType = "MHR"
	MHW GameType = "MHW"
)

type GameType string

func IsGameType(s string) bool {
	switch s {
	case string(MHR):
		return true
	case string(MHW):
		return true

	default:
		return false
	}
}

type BackupUploadModel struct {
	Id         string
	Size       int64
	Game       GameType
	UploadedAt time.Time
}
