package localization

import (
	"github.com/Haato3o/poogie/pkg/aws"
	"github.com/Haato3o/poogie/pkg/crypto"
	"github.com/Haato3o/poogie/pkg/memcache"
	"github.com/Haato3o/poogie/pkg/server"
	"github.com/gin-gonic/gin"
	"time"
)

type Handler struct{}

func (h Handler) Load(router *gin.RouterGroup, server *server.Server) error {
	service := NewService(
		aws.New(server.Config, "localization/", ".xml"),
		crypto.NewHashService(server.Config.HashSalt),
		memcache.New(5*time.Minute),
	)

	controller := NewController(service)

	router.GET("/localization/checksum", controller.GetLocalizationsHandler)

	return nil
}

func (h Handler) GetVersion() int {
	return server.V1
}

func (h Handler) GetName() string {
	return "LocalizationHandler"
}

func New() server.IRegisterableService {
	return &Handler{}
}
