package monitoring

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
)

// nullMonitoring is a no-op implementation of Monitoring that is used by default
type nullMonitoring struct{}

func (n *nullMonitoring) WrapHTTP(handler http.Handler) http.Handler {
	return handler
}

func (n *nullMonitoring) WrapHTTPTransport(r http.RoundTripper) http.RoundTripper {
	return r
}

func (n *nullMonitoring) GRPCUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return nil
}

func (n *nullMonitoring) GRPCUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return nil
}

func (n *nullMonitoring) AddField(ctx context.Context, key string, value interface{}) {}

func (n *nullMonitoring) AddFieldToTrace(ctx context.Context, key string, value interface{}) {}

func (n *nullMonitoring) Close() error {
	return nil
}
