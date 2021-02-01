package monitoring

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
)

var Default Monitoring = &nullMonitoring{}

type Monitoring interface {
	WrapHTTP(handler http.Handler) http.Handler
	WrapHTTPTransport(r http.RoundTripper) http.RoundTripper
	GRPCUnaryClientInterceptor() grpc.UnaryClientInterceptor
	GRPCUnaryServerInterceptor() grpc.UnaryServerInterceptor
	AddField(ctx context.Context, key string, value interface{})
	AddFieldToTrace(ctx context.Context, key string, value interface{})
	Close() error
}

func Close() error {
	return Default.Close()
}
