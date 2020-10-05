package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fullstorydev/grpcui/standalone"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid/rpcgrid"
	"github.com/theothertomelliott/tic-tac-toverengineered/internal/turncontroller"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/turn/inmemoryturns"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/turn/rpcturn"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/win/rpcchecker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func getCurrentTurnServerTarget() string {
	if serverTarget := os.Getenv("CURRENT_TURN_SERVER_TARGET"); serverTarget != "" {
		return serverTarget
	}
	return "localhost:8084"
}

func getGridServerTarget() string {
	if serverTarget := os.Getenv("GRID_SERVER_TARGET"); serverTarget != "" {
		return serverTarget
	}
	return "localhost:8086"
}

func getCheckerServerTarget() string {
	if serverTarget := os.Getenv("CHECKER_SERVER_TARGET"); serverTarget != "" {
		return serverTarget
	}
	return "localhost:8088"
}

func main() {
	port := 8080
	grpcuiPort := 8081
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	g, err := rpcgrid.ConnectGrid(getGridServerTarget())
	if err != nil {
		log.Fatalf("could not connect to grid server: %v", err)
	}
	checker, err := rpcchecker.ConnectChecker(getCheckerServerTarget())
	if err != nil {
		log.Fatalf("could not connect to win checker server: %v", err)
	}
	ct, err := rpcturn.ConnectCurrent(getCurrentTurnServerTarget())
	if err != nil {
		log.Fatalf("could not connect to current turn server: %v", err)
	}
	controllerBackend := inmemoryturns.New(ct, g, checker)
	rpcturn.RegisterControllerServer(grpcServer, turncontroller.NewServer(controllerBackend))
	log.Printf("gRPC listening on port :%v", port)

	// we need the reflection service, to power the UI
	reflection.Register(grpcServer)

	var done = make(chan struct{})

	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err := startGrpcUI(port, grpcuiPort)
		if err != nil {
			log.Printf("Failed to start gRPCUI: %v", err)
		}
	}()

	<-done
}

func startGrpcUI(port, grpcuiPort int) error {
	// Create a connection to local gRPC
	serverAddr := fmt.Sprintf("127.0.0.1:%d", port)
	cc, err := grpc.Dial(serverAddr, grpc.WithInsecure())
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
