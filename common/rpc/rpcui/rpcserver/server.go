package rpcserver

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// New creates an RPC server that will listen on the specified port
func New(port int) *Server {
	return &Server{
		grpcServer: grpc.NewServer(),
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
