package tracing

type ITracingTransaction interface {
	AddProperty(key string, value interface{})
}
