package version

import (
	"github.com/Haato3o/poogie/pkg/aws"
	"github.com/Haato3o/poogie/pkg/server"
	"github.com/gin-gonic/gin"
)

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
	service := NewService(
		aws.New(server.Config, "Releases/", ".zip"),
		aws.New(server.Config, "Beta/", ".zip"),
		server.Database.GetSupporterRepository(),
		server.Database.GetPatchRepository(),
	)

	controller := VersionController{
		service: service,
	}

	router.GET("/version", controller.GetLatestVersion)
	router.GET("/version/latest", controller.GetLatestBinary)
	router.GET("/version/:version", controller.GetBinaryByVersion)
	router.GET("/patch/notes", controller.GetAllPatchNotes)

	return nil
}

func New() server.IRegisterableService {
	return &VersionHandler{}
}
