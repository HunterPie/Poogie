package middlewares

import (
	"github.com/Haato3o/poogie/pkg/log"
	"github.com/gin-gonic/gin"
)

func LogRequest(ctx *gin.Context) {
	log.Info(ctx.Request.URL.Path)
}
