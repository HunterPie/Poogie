package session

import (
	"github.com/Haato3o/poogie/pkg/http"
	"github.com/Haato3o/poogie/pkg/server"
	"github.com/gin-gonic/gin"
)

type SessionHandler struct{}

type SessionResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

var sessionResponse = SessionResponse{
	Ok:      true,
	Message: "route has been deprecated",
}

// GetName implements server.IRegisterableService
func (*SessionHandler) GetName() string {
	return "SessionHandler"
}

// GetVersion implements server.IRegisterableService
func (*SessionHandler) GetVersion() int {
	return server.V1
}

// Load implements server.IRegisterableService
func (*SessionHandler) Load(router *gin.RouterGroup, server *server.Server) error {

	// DEPRECATED ROUTES
	router.GET("/session", func(ctx *gin.Context) { http.Ok(ctx, sessionResponse) })
	router.POST("/session/end", func(ctx *gin.Context) { http.Ok(ctx, sessionResponse) })

	router.POST("/login")

	return nil
}

func New() server.IRegisterableService {
	return &SessionHandler{}
}
