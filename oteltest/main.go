package main

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/exporters/stdout"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	exporter, err := stdout.NewExporter(
		stdout.WithPrettyPrint(),
	)
	if err != nil {
		log.Fatalf("failed to initialize stdout export pipeline: %v", err)
	}

	ctx := context.Background()
	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(bsp))

	// Handle this error in a sensible manner where possible
	defer func() { _ = tp.Shutdown(ctx) }()

	pusher := controller.New(
		processor.New(
			simple.NewWithExactDistribution(),
			exporter,
		),
		controller.WithPusher(exporter),
		controller.WithCollectPeriod(5*time.Second),
	)

	err = pusher.Start(ctx)
	if err != nil {
		log.Fatalf("failed to initialize metric controller: %v", err)
	}

	// Handle this error in a sensible manner where possible
	defer func() { _ = pusher.Stop(ctx) }()

	otel.SetTracerProvider(tp)
	otel.SetMeterProvider(pusher.MeterProvider())
	propagator := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
	otel.SetTextMapPropagator(propagator)

	fooKey := label.Key("ex.com/foo")
	barKey := label.Key("ex.com/bar")
	lemonsKey := label.Key("ex.com/lemons")
	anotherKey := label.Key("ex.com/another")

	commonLabels := []label.KeyValue{lemonsKey.Int(10), label.String("A", "1"), label.String("B", "2"), label.String("C", "3")}

	meter := otel.Meter("ex.com/basic")

	observerCallback := func(_ context.Context, result metric.Float64ObserverResult) {
		result.Observe(1, commonLabels...)
	}
	_ = metric.Must(meter).NewFloat64ValueObserver("ex.com.one", observerCallback,
		metric.WithDescription("A ValueObserver set to 1.0"),
	)

	valueRecorder := metric.Must(meter).NewFloat64ValueRecorder("ex.com.two")

	boundRecorder := valueRecorder.Bind(commonLabels...)
	defer boundRecorder.Unbind()

	tracer := otel.Tracer("ex.com/basic")
	ctx = baggage.ContextWithValues(ctx,
		fooKey.String("foo1"),
		barKey.String("bar1"),
	)

	func(ctx context.Context) {
		var span trace.Span
		ctx, span = tracer.Start(ctx, "operation")
		defer span.End()

		span.AddEvent("Nice operation!", trace.WithAttributes(label.Int("bogons", 100)))
		span.SetAttributes(anotherKey.String("yes"))

		meter.RecordBatch(
			// Note: call-site variables added as context Entries:
			baggage.ContextWithValues(ctx, anotherKey.String("xyz")),
			commonLabels,

			valueRecorder.Measurement(2.0),
		)

		func(ctx context.Context) {
			var span trace.Span
			ctx, span = tracer.Start(ctx, "Sub operation...")
			defer span.End()

			span.SetAttributes(lemonsKey.String("five"))
			span.AddEvent("Sub span event")
			boundRecorder.Record(ctx, 1.3)
		}(ctx)
	}(ctx)
}
