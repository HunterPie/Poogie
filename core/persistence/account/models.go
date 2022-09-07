package account

import "time"

type AccountBadgesModel struct {
	Id        string
	CreatedAt time.Time
}

type HuntStatisticsSummaryModel struct {
	Id        string
	CreatedAt time.Time
}

type AccountModel struct {
	Id                         string
	Username                   string
	Password                   string
	Email                      string
	ClientId                   string
	Experience                 int64
	Rating                     int64
	AvatarUri                  string
	Badges                     []AccountBadgesModel
	HuntStatisticsSummaryModel []HuntStatisticsSummaryModel
	IsSupporter                bool
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
	LastSessionAt              time.Time
	IsArchived                 bool
}
