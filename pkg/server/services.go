package server

import "github.com/gin-gonic/gin"

const (
	NO_VERSION int = -1
	V1         int = iota
	V2
)

type IRegisterableService interface {
	Load(router *gin.RouterGroup, server *Server) error
	GetVersion() int
	GetName() string
}
