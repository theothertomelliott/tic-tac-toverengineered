package rpcui

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fullstorydev/grpcui/standalone"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

// Start configures and starts a grpcui server
// connecting to localhost on the specified port.
// This call will block while the server is running
// successfully.
func Start(port, grpcuiPort int) error {
	// Create a connection to local gRPC
	serverAddr := fmt.Sprintf("127.0.0.1:%d", port)
	cc, err := grpc.Dial(
		serverAddr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	if err != nil {
		return fmt.Errorf("failed to connect to localhost: %w", err)
	}

	// Create the grpcui handler
	target := fmt.Sprintf("%s:%d", filepath.Base(os.Args[0]), port)
	h, err := standalone.HandlerViaReflection(context.Background(), cc, target)
	if err != nil {
		return fmt.Errorf("failed to create handler for local server %q: %w", target, err)
	}

	// Add to an http server
	serveMux := http.NewServeMux()
	serveMux.Handle("/", h)
	log.Printf("grpcui listening on port :%v", grpcuiPort)
	http.ListenAndServe(fmt.Sprintf(":%v", grpcuiPort), serveMux)

	return nil
}
