package opentelemetry

import (
	"context"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

const (
	ENV_OTEL_JAEGER_ENDPOINT = "OTEL_JAEGER_ENDPOINT"
)

// Setup configures opentelemetry for the named service
func Setup(serviceName string) (func() error, error) {
	ctx := context.Background()

	exp, err := newExporter(ctx)
	if err != nil {
		return nil, err
	}

	// Create a new tracer provider with a batch span processor and the given exporter.
	tp := newTraceProvider(exp, serviceName)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return func() error { return tp.Shutdown(ctx) }, nil
}

// newExporter returns an opentelemetry exporter configured according to environment variables
func newExporter(context.Context) (sdktrace.SpanExporter, error) {
	if os.Getenv(ENV_OTEL_JAEGER_ENDPOINT) != "" {
		url := os.Getenv(ENV_OTEL_JAEGER_ENDPOINT)
		return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	}

	// Default to stdout logging
	w := os.Stdout
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
	)
}

func newTraceProvider(exp sdktrace.SpanExporter, serviceName string) *sdktrace.TracerProvider {
	// Ensure default SDK resources and the required service name are set.
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		),
	)

	if err != nil {
		panic(err)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)
}
