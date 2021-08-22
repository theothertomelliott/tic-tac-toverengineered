package otelmonitoring

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	"google.golang.org/grpc"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var _ monitoring.Monitoring = &Monitoring{}

func New(componentName string) (monitoring.Monitoring, error) {
	ctx := context.Background()

	// Set up a trace exporter
	ex, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint("otel-collector:55680"),
		otlptracegrpc.WithDialOption(grpc.WithBlock()),
	)
	if err != nil {
		return nil, err
	}

	resources, err := resource.New(
		ctx,
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String(componentName),
			semconv.ServiceVersionKey.String(version.Version),
			attribute.String("library.language", "go"),
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
