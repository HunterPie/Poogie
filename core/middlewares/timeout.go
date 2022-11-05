package middlewares

import (
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(timeoutFunc gin.HandlerFunc) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(5*time.Minute),
		timeout.WithHandler(func(ctx *gin.Context) {
			ctx.Next()
		}),
		timeout.WithResponse(timeoutFunc),
	)
}
