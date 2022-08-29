package tracing

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type NewRelicTransaction struct {
	*newrelic.Transaction
}

func (t *NewRelicTransaction) AddProperty(key string, value interface{}) {
	t.AddAttribute(key, value)
}

func FromContext(ctx context.Context) ITracingTransaction {
	txn := newrelic.FromContext(ctx)
	return &NewRelicTransaction{txn}
}
