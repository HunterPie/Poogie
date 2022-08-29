package middlewares

import (
	"github.com/Haato3o/poogie/core/tracing"
	"github.com/gin-gonic/gin"
)

func LoggingMiddleware(ctx *gin.Context) {
	txn = tracing.FromContext(ctx)

	// TODO: Implement this
}
