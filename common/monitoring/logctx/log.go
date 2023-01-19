package logctx

import (
	"context"
	"fmt"
	"log"

	"go.opentelemetry.io/otel/trace"
)

func Println(ctx context.Context, args ...interface{}) {
	message := fmt.Sprint(args...)

	span := trace.SpanFromContext(ctx)
	traceid := "<none>"
	if span != nil {
		traceid = span.SpanContext().TraceID().String()
	}

	log.Printf("traceid=%v > %v", traceid, message)

	if span != nil {
		span.AddEvent(message)
	}
}

func Printf(ctx context.Context, format string, args ...interface{}) {
	Println(ctx, fmt.Sprintf(format, args...))
}
