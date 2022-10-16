package config

import (
	"fmt"
	"time"

	"github.com/Haato3o/poogie/core/services"
	"github.com/Haato3o/poogie/metadata"
)

var DeployEmbed = services.DiscordEmbed{
	Title:       "Poogie - Deploy",
	Description: "Poogie has been deployed",
	Color:       6881177,
	Fields: []services.DiscordEmbedField{
		{
			Name:  "Version",
			Value: metadata.Version,
		},
		{
			Name:  "Deployed At",
			Value: fmt.Sprintf("<t:%d:R>", time.Now().Unix()),
		},
	},
}

var DiedEmbed = services.DiscordEmbed{
	Title:       "Poogie - Warning",
	Description: "Poogie instance has been taken down",
	Color:       16728703,
	Fields: []services.DiscordEmbedField{
		{
			Name:  "Version",
			Value: metadata.Version,
		},
		{
			Name:  "Timestamp",
			Value: fmt.Sprintf("<t:%d:R>", time.Now().Unix()),
		},
	},
}
