package otelmonitoring

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpgrpc"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var _ monitoring.Monitoring = &Monitoring{}

func New() (monitoring.Monitoring, error) {
	ctx := context.Background()

	ex, err := otlp.NewExporter(ctx, otlpgrpc.NewDriver(
		otlpgrpc.WithEndpoint("otel-collector:55680"),
		otlpgrpc.WithInsecure(),
	))
	if err != nil {
		return nil, err
	}

	bsp := sdktrace.NewBatchSpanProcessor(ex)
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
