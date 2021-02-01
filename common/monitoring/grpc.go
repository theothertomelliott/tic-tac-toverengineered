package monitoring

import (
	"google.golang.org/grpc"
)

func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return Default.GRPCUnaryClientInterceptor()
}

// UnaryServerInterceptor starts a beeline span for each grpc call.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return Default.GRPCUnaryServerInterceptor()
}
