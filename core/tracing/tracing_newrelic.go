package tracing

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type NewRelicTransaction struct {
	*newrelic.Transaction
}

type NewRelicSegment struct {
	*newrelic.Segment
}

func (s *NewRelicSegment) End() {
	s.Segment.End()
}

// StartSegment implements ITracingTransaction
func (t *NewRelicTransaction) StartSegment(name string) ITracingSegment {
	segment := newrelic.Segment{
		Name:      name,
		StartTime: t.StartSegmentNow(),
	}

	return &NewRelicSegment{&segment}
}

// AddProperty implements ITracingTransaction
func (t *NewRelicTransaction) AddProperty(key string, value interface{}) {
	t.AddAttribute(key, value)
}

func FromContext(ctx context.Context) ITracingTransaction {
	txn := newrelic.FromContext(ctx)
	return &NewRelicTransaction{txn}
}
