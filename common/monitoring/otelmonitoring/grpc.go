package otelmonitoring

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

const TraceHeader string = "x-trace-headers"

func (m *Monitoring) GRPCUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return otelgrpc.UnaryClientInterceptor()
}

// GRPCUnaryServerInterceptor starts a beeline span for each grpc call.
func (m *Monitoring) GRPCUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return otelgrpc.UnaryServerInterceptor()
}
