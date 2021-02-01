package logmonitoring

import (
	"context"
	"log"

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
	log.Println("creating gRPC client interceptor")
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		log.Println("gRPC client call: ", method)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// GRPCUnaryServerInterceptor starts a beeline span for each grpc call.
func (m *Monitoring) GRPCUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	log.Println("creating gRPC server interceptor")
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Println("gRPC server call: ", info.FullMethod)
		return handler(ctx, req)
	}
}
