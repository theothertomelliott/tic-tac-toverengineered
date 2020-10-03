package main

import (
	"fmt"
	"log"
	"net"
	"os"

	space "github.com/theothertomelliott/tic-tac-toverengineered/space/internal"
	"github.com/theothertomelliott/tic-tac-toverengineered/space/pkg/rpcspace"
	"github.com/theothertomelliott/tic-tac-toverengineered/space/pkg/spaceinmemory"
	"google.golang.org/grpc"
)

func getPort() string {
	if serverTarget := os.Getenv("PORT"); serverTarget != "" {
		return serverTarget
	}
	return "8080"
}

func main() {
	port := getPort()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	spaceBackend := spaceinmemory.New()
	rpcspace.RegisterSpaceServer(grpcServer, space.NewServer(spaceBackend))
	log.Printf("gRPC listening on port :%v", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
