package tracing

import "context"

type TracingEngine interface {
	FromContext(ctx *context.Context) *TracingTransaction
}

type TracingTransaction interface {
	AddProperty(key string, value interface{})
}
