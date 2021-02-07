package otelmonitoring

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

// blacklistedMethods are methods that should not be traced such as the
// gRPC healthcheck service.
//
// Currently all of these are Unary methods so we do not check this list in the
// StreamServerInterceptor.
var blacklistedMethods = map[string]struct{}{
	"/grpc.health.v1.Health/Check": {},
}

const TraceHeader string = "x-trace-headers"

func (m *Monitoring) GRPCUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return otelgrpc.UnaryClientInterceptor()
}

// GRPCUnaryServerInterceptor starts a beeline span for each grpc call.
func (m *Monitoring) GRPCUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return otelgrpc.UnaryServerInterceptor()
}
