package honeycomb

import (
	"context"
	"strings"

	beeline "github.com/honeycombio/beeline-go"
	"github.com/honeycombio/beeline-go/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx, sp := beeline.StartSpan(ctx, method)
		defer sp.Send()
		if sp != nil {
			pc := sp.SerializeHeaders()
			md := metadata.New(map[string]string{TraceHeader: pc})
			ctx = metadata.NewOutgoingContext(ctx, md)
		}

		err := invoker(ctx, method, req, reply, cc, opts...)
		beeline.AddField(ctx, "grpc_code", status.Code(err).String())
		return err
	}
}

// GRPCUnaryServerInterceptor starts a beeline span for each grpc call.
func (m *Monitoring) GRPCUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if _, exists := blacklistedMethods[info.FullMethod]; exists {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.DataLoss, "failed to get metadata")
		}

		xrid := md[TraceHeader]
		serializedHeaders := ""
		if len(xrid) != 0 {
			if strings.Trim(xrid[0], " ") == "" {
				return nil, status.Errorf(codes.InvalidArgument, "empty 'x-request-id' header")
			}
			serializedHeaders = xrid[0]
		}
		var tr *trace.Trace
		ctx, tr = trace.NewTrace(ctx, serializedHeaders)
		defer tr.Send()

		resp, err := handler(ctx, req)

		beeline.AddField(ctx, "grpc_code", status.Code(err).String())

		return resp, err
	}
}
