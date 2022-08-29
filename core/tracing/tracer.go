package tracing

import "context"

type ITracingEngine interface {
	FromContext(ctx context.Context) ITracingTransaction
}

type ITracingTransaction interface {
	AddProperty(key string, value interface{})
}
