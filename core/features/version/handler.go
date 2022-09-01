package version

import (
	"github.com/Haato3o/poogie/pkg/aws"
	"github.com/Haato3o/poogie/pkg/server"
	"github.com/gin-gonic/gin"
)

const SCOPE = "/version"

type VersionHandler struct{}

// GetName implements server.IRegisterableService
func (*VersionHandler) GetName() string {
	return "VersionHandler"
}

// GetVersion implements server.IRegisterableService
func (*VersionHandler) GetVersion() int {
	return server.V1
}

// Load implements server.IRegisterableService
func (*VersionHandler) Load(router *gin.RouterGroup, server *server.Server) error {
	service := VersionService{
		bucket:              aws.New(server.Config, "Releases/", ".zip"),
		alphaBucket:         aws.New(server.Config, "Beta/", ".zip"),
		supporterRepository: server.Database.GetSupporterRepository(),
	}

	controller := VersionController{
		service: &service,
	}

	router.GET(SCOPE, controller.GetLatestVersion)
	router.GET(subScope("/:version"), controller.GetBinaryByVersion)

	return nil
}

func New() server.IRegisterableService {
	return &VersionHandler{}
}

func subScope(path string) string {
	return SCOPE + path
}
