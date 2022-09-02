package http

import (
	"net/http"

	"github.com/Haato3o/poogie/core/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Endpoint interface {
	Router(*gin.RouterGroup)
}

type HttpServer struct {
	server *http.Server
	Router *gin.Engine
}

func New(address string) *HttpServer {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(cors.Default())
	router.Use(middlewares.TransactionMiddleware)

	server := &http.Server{
		Addr:    address,
		Handler: router,
	}

	return &HttpServer{
		server: server,
		Router: router,
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
