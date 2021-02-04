package honeycomb

import (
	"context"

	"github.com/honeycombio/beeline-go"
	"github.com/honeycombio/beeline-go/trace"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
)

func (m *Monitoring) StartSpan(ctx context.Context, name string) (context.Context, monitoring.Span) {
	ctx, sp := beeline.StartSpan(ctx, name)
	return ctx, &span{span: sp}
}

type span struct {
	span *trace.Span
}

func (s *span) Finish() error {
	s.span.Send()
	return nil
}
