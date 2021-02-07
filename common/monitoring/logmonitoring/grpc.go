package logmonitoring

import (
	"context"
	"encoding/json"
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
		err := invoker(ctx, method, req, reply, cc, opts...)
		var errMsg string
		if err != nil {
			errMsg = err.Error()
		}
		log.Printf("gRPC out: method=%q, request=%q, reply=%q, err=%q", method, renderJSON(req), renderJSON(reply), errMsg)
		return err
	}
}

// GRPCUnaryServerInterceptor starts a beeline span for each grpc call.
func (m *Monitoring) GRPCUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	log.Println("creating gRPC server interceptor")
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Printf("gRPC in: method=%q, req=%q", info.FullMethod, renderJSON(req))
		return handler(ctx, req)
	}
}

func renderJSON(in interface{}) string {
	data, err := json.Marshal(in)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
