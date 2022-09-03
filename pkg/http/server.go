package http

import (
	"net/http"

	"github.com/Haato3o/poogie/core/config"
	"github.com/Haato3o/poogie/core/middlewares"
	"github.com/Haato3o/poogie/core/tracing"
	"github.com/Haato3o/poogie/pkg/newrelic"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Endpoint interface {
	Router(*gin.RouterGroup)
}

type HttpServer struct {
	server        *http.Server
	Router        *gin.Engine
	tracingEngine tracing.ITracingEngine
}

func New(config *config.ApiConfiguration, isMonitoringEnabled bool) *HttpServer {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(cors.Default())
	router.Use(middlewares.TransactionMiddleware)

	var tracer tracing.ITracingEngine
	if isMonitoringEnabled {
		tracer = newrelic.New(
			config.AppEnv,
			config.NewRelicLicenseKey,
			!config.Debug,
		)
	}

	server := &http.Server{
		Addr:    config.HttpAddress,
		Handler: router,
	}

	return &HttpServer{
		server:        server,
		Router:        router,
		tracingEngine: tracer,
	}
}

func (s *HttpServer) Group(path string) *gin.RouterGroup {
	return s.Router.Group(path)
}

func (s *HttpServer) Start() error {
	s.Router.NoRoute(NotFoundHandler)
	s.server.Handler = s.Router
	return s.server.ListenAndServe()
}

func (s *HttpServer) Stop() {
	s.server.Close()
}
