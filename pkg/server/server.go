package server

import (
	"fmt"
	"github.com/Haato3o/poogie/pkg/crypto"
	"github.com/Haato3o/poogie/pkg/jwt"

	"github.com/Haato3o/poogie/core/config"
	"github.com/Haato3o/poogie/core/middlewares"
	"github.com/Haato3o/poogie/core/persistence/database"
	"github.com/Haato3o/poogie/pkg/http"
	"github.com/Haato3o/poogie/pkg/log"
	"github.com/Haato3o/poogie/pkg/mongodb"
)

type Server struct {
	Config     *config.ApiConfiguration
	HttpServer *http.HttpServer
	Database   database.IDatabase
	quit       chan struct{}
}

func New(config *config.ApiConfiguration) (*Server, error) {
	isMonitoringEnabled := config.NewRelicLicenseKey != ""
	server := http.New(config, isMonitoringEnabled)
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

func (s *Server) Load(services map[int][]IRegisterableService) error {
	userIdTransform := middlewares.NewUserTransformMiddleware(
		jwt.New(s.Config.JwtKey),
		s.Database.GetSessionRepository(),
		crypto.NewHashService(
			s.Config.HashSalt,
		),
	)
	supporterTransform := middlewares.NewSupporterTransformMiddleware(s.Database)

	for version, servicesList := range services {
		group := ""
		if version > NO_VERSION {
			group = fmt.Sprintf("v%d", version)
		}

		router := s.HttpServer.Router.Group(group)

		if version > NO_VERSION {
			router.Use(userIdTransform.TokenToUserIdTransform)
			router.Use(supporterTransform.TransformSupporterToken)
		}

		router.Use(middlewares.LogRequest)
		router.Use(middlewares.TransactionMiddleware)

		for _, service := range servicesList {
			err := service.Load(router, s)

			log.Info(fmt.Sprintf("Registering handler %s:%s", service.GetName(), group))

			if err != nil {
				return err
			}
		}
	}

	return nil
}
