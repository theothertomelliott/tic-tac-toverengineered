package rpcserver

import (
	"fmt"
	"net"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/rakyll/goutil/pprofutil"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/filters"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// New creates an RPC server that will listen on the specified port
// A default grpc health server (google.golang.org/grpc/health) will also be hosted on this port.
func New(port int) *Server {
	return NewWithHealthServer(port, nil)
}

// NewWithHealthServer creates an RPC server that will listen on the specified port
// The provided grpc health server (google.golang.org/grpc/health) will also be hosted on this port.
func NewWithHealthServer(port int, healthServer *health.Server) *Server {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			pprofutil.UnaryServerInterceptor(),
			otelgrpc.UnaryServerInterceptor(
				// Do not trace health checks
				otelgrpc.WithInterceptorFilter(
					filters.Not(
						filters.HealthCheck(),
					),
				),
			),
			grpc_prometheus.UnaryServerInterceptor,
		),
	)

	// Add health check for all rpc servers
	if healthServer == nil {
		healthServer = health.NewServer()
	}
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	return &Server{
		grpcServer: grpcServer,
		port:       port,
	}
}

// Server defines an RPC server
type Server struct {
	grpcServer *grpc.Server
	port       int
}

// Serve begins serving RPC endpoints on the configured port
func (s *Server) Serve() error {
	// enable reflection for grpcui
	reflection.Register(s.grpcServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("could not listen on port %d: %w", s.port, err)
	}
	return s.grpcServer.Serve(lis)
}

// GRPC gets the gRPC server, to be used when registering handlers
func (s *Server) GRPC() *grpc.Server {
	return s.grpcServer
}
