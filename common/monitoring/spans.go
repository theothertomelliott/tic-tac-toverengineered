package monitoring

import "context"

type Span interface {
	Finish() error
}

func StartSpan(ctx context.Context, name string) (context.Context, Span) {
	return Default.StartSpan(ctx, name)
}
