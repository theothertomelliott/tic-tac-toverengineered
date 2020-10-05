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
	gridserver "github.com/theothertomelliott/tic-tac-toverengineered/grid/internal/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid/rpcgrid"
	space "github.com/theothertomelliott/tic-tac-toverengineered/space/pkg"
	"github.com/theothertomelliott/tic-tac-toverengineered/space/pkg/rpcspace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 8080
	grpcuiPort := 8081
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	var spaces [][]space.Space
	for i := 0; i < 3; i++ {
		var row []space.Space
		for j := 0; j < 3; j++ {
			c, err := rpcspace.ConnectSpace(fmt.Sprintf("space-%d-%d:80", i, j))
			if err != nil {
				log.Fatalf("space (%d,%d): %v", i, j, err)
			}
			row = append(row, c)
		}
		spaces = append(spaces, row)
	}
	gridBackend, _ := grid.New(spaces)

	rpcgrid.RegisterGridServer(grpcServer, gridserver.NewServer(gridBackend))
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
