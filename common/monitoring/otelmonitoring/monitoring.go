package otelmonitoring

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var _ monitoring.Monitoring = &Monitoring{}

func New() (monitoring.Monitoring, error) {
	exporter, err := stdout.NewExporter()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(bsp))

	otel.SetTracerProvider(tp)

	return &Monitoring{
		ctx: ctx,
		tp:  tp,
	}, nil
}

type Monitoring struct {
	ctx context.Context
	tp  *sdktrace.TracerProvider
}

func (m *Monitoring) Close() error {
	return m.tp.Shutdown(m.ctx)
}
