package middlewares

import (
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(timeoutFunc gin.HandlerFunc) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(60*time.Second),
		timeout.WithHandler(func(ctx *gin.Context) {
			ctx.Next()
		}),
		timeout.WithResponse(timeoutFunc),
	)
}
