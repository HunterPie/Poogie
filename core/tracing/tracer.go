package tracing

import "github.com/gin-gonic/gin"

type ITracingTransaction interface {
	AddProperty(key string, value interface{})
}

type ITracingEngine interface {
	SetupTracingMiddleware(router *gin.Engine)
}
