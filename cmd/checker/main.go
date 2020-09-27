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
	checkerserver "github.com/theothertomelliott/tic-tac-toverengineered/internal/checker"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/grid/rpcgrid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/win/gridchecker"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/win/rpcchecker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func getGridServerTarget() string {
	if serverTarget := os.Getenv("GRID_SERVER_TARGET"); serverTarget != "" {
		return serverTarget
	}
	return "localhost:8086"
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
	checkerBackend := gridchecker.New(g)
	rpcchecker.RegisterCheckerServer(grpcServer, checkerserver.NewServer(checkerBackend))
	log.Printf("gRPC listening on port :%v", port)
	go grpcServer.Serve(lis)

	// we need the reflection service, to power the UI
	reflection.Register(grpcServer)

	// Create a connection to local gRPC
	cc, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to create client to local server: %v", err)
	}

	// Create the grpcui handler
	target := fmt.Sprintf("%s:%d", filepath.Base(os.Args[0]), port)
	h, err := standalone.HandlerViaReflection(context.Background(), cc, target)
	if err != nil {
		log.Fatalf("failed to create client to local server: %v", err)
	}

	// Add to an http server
	serveMux := http.NewServeMux()
	serveMux.Handle("/", h)
	log.Printf("grpcui listening on port :%v", grpcuiPort)
	http.ListenAndServe(fmt.Sprintf(":%v", grpcuiPort), serveMux)
}
