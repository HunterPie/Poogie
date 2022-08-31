package server

import (
	"github.com/Haato3o/poogie/core/config"
	"github.com/Haato3o/poogie/core/domain/persistence/database"
	"github.com/Haato3o/poogie/core/middlewares"
	"github.com/Haato3o/poogie/pkg/http"
	"github.com/Haato3o/poogie/pkg/mongodb"
)

type Server struct {
	Config     *config.ApiConfiguration
	HttpServer *http.HttpServer
	Database   database.IDatabase

	quit chan struct{}
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

func (s *Server) Start() {
	go s.HttpServer.Start()
	<-s.quit
}

func (s *Server) Stop() {
	s.HttpServer.Stop()
	s.quit <- struct{}{}
}

func (s *Server) Load(group string, services ...IRegisterableService) error {
	router := s.HttpServer.Router.Group(group)
	router.Use(middlewares.TransactionMiddleware)

	for _, svc := range services {
		err := svc.Load(router, s)

		if err != nil {
			return err
		}
	}

	return nil
}
