package tracing

import "github.com/gin-gonic/gin"

type ITracingSegment interface {
	End()
}

type ITracingTransaction interface {
	AddProperty(key string, value interface{})
	StartSegment(name string) ITracingSegment
}

type ITracingEngine interface {
	SetupTracingMiddleware(router *gin.Engine)
}
