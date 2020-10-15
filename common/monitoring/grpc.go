package monitoring

import (
	"context"
	"log"
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

func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// Anything linked to this variable will transmit request headers.
		sp := trace.GetSpanFromContext(ctx)
		if sp != nil {
			pc := sp.SerializeHeaders()
			md := metadata.New(map[string]string{TraceHeader: pc})
			ctx = metadata.NewOutgoingContext(ctx, md)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// UnaryServerInterceptor starts a beeline span for each grpc call.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Printf("Intercepting request: %v", info.FullMethod)
		if _, exists := blacklistedMethods[info.FullMethod]; exists {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.DataLoss, "failed to get metadata")
		}

		var (
			tr *trace.Trace
			sp *trace.Span
		)

		sp = trace.GetSpanFromContext(ctx)
		if sp == nil {

			xrid := md[TraceHeader]
			if len(xrid) == 0 {
				ctx, sp = beeline.StartSpan(ctx, info.FullMethod)
				tr = sp.GetTrace()
			} else {
				if strings.Trim(xrid[0], " ") == "" {
					return nil, status.Errorf(codes.InvalidArgument, "empty 'x-request-id' header")
				}
				ctx, tr = trace.NewTraceFromSerializedHeaders(ctx, xrid[0])
				sp = tr.GetRootSpan()
				// ctx, sp = sp.CreateChild(ctx)
				// ctx, sp = beeline.StartSpan(ctx, info.FullMethod)
			}
		} else {
			tr = sp.GetTrace()
		}
		sp.AddField("name", info.FullMethod)
		log.Println(tr)
		defer tr.Send()

		resp, err := handler(ctx, req)

		addRPCTrace(ctx, err)

		return resp, err
	}
}

func addRPCTrace(ctx context.Context, err error) {
	beeline.AddFieldToTrace(ctx, "grpc_code", status.Code(err).String())
	// if requestID := middleware.RequestID(ctx); requestID != "" {
	// 	beeline.AddFieldToTrace(ctx, "request_id", requestID)
	// }
	// if caller := middleware.Caller(ctx); caller != "" {
	// 	beeline.AddField(ctx, "caller", caller)
	// }
}
