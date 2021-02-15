package otelmonitoring

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpgrpc"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var _ monitoring.Monitoring = &Monitoring{}

func New(componentName string) (monitoring.Monitoring, error) {
	ctx := context.Background()

	ex, err := otlp.NewExporter(ctx, otlpgrpc.NewDriver(
		otlpgrpc.WithEndpoint("otel-collector:55680"),
		otlpgrpc.WithInsecure(),
	))
	if err != nil {
		return nil, err
	}

	resources, err := resource.New(
		ctx,
		resource.WithAttributes(
			label.String("service.name", componentName),
			label.String("service.version", version.Version),
			label.String("library.language", "go"),
		),
	)
	if err != nil {
		return nil, err
	}

	bsp := sdktrace.NewBatchSpanProcessor(ex)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(bsp),
		sdktrace.WithResource(resources),
	)

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
