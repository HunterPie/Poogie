package server

import (
	"github.com/Haato3o/poogie/core/config"
	"github.com/Haato3o/poogie/pkg/http"
	"github.com/Haato3o/poogie/pkg/mongodb"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Config     *config.ApiConfiguration
	HttpServer *http.HttpServer
	Database   *mongo.Client

	quit chan struct{}
}

type RegisterableService interface {
	Load(router *gin.RouterGroup, server *Server) error
}

func New(config *config.ApiConfiguration) (*Server, error) {
	server := http.New(config.HttpAddress)

	isMonitoringEnabled := config.NewRelicLicenseKey != ""
	database, err := mongodb.New(config.DatabaseUri, config.DatabaseName, isMonitoringEnabled)

	if err != nil {
		return nil, err
	}

	return &Server{
		Config:     config,
		HttpServer: server,
		Database:   database,
	}, nil
}
