package monitoring

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/grpc"
)

// nullMonitoring is a no-op implementation of Monitoring that is used by default
type nullMonitoring struct{}

func (n *nullMonitoring) WrapHTTP(handler http.Handler) http.Handler {
	log.Println("WrapHTTP: monitoring has not been initialized")
	return handler
}

func (n *nullMonitoring) WrapHTTPTransport(r http.RoundTripper) http.RoundTripper {
	log.Println("WrapHTTPTransport: monitoring has not been initialized")
	return r
}

func (n *nullMonitoring) GRPCUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	log.Println("GRPCUnaryClientInterceptor: monitoring has not been initialized")
	return nil
}

func (n *nullMonitoring) GRPCUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	log.Println("GRPCUnaryServerInterceptor: monitoring has not been initialized")
	return nil
}

func (n *nullMonitoring) AddField(ctx context.Context, key string, value interface{}) {
	log.Println("AddField: monitoring has not been initialized")
}

func (n *nullMonitoring) AddFieldToTrace(ctx context.Context, key string, value interface{}) {
	log.Println("AddFieldToTrace: monitoring has not been initialized")
}

func (n *nullMonitoring) Close() error {
	return nil
}
