package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpcui"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid/rpcgrid"
	checkerserver "github.com/theothertomelliott/tic-tac-toverengineered/internal/checker"
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

	// we need the reflection service, to power the UI
	reflection.Register(grpcServer)
	log.Printf("gRPC listening on port :%v", port)

	var done = make(chan struct{})
	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		err := rpcui.Start(port, grpcuiPort)
		if err != nil {
			log.Printf("Failed to start gRPCUI: %v", err)
		}
	}()
	<-done
}
