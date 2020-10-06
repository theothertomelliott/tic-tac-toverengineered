package main

import (
	"fmt"
	"log"
	"net"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/internal/gamerepo"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game/inmemoryrepository"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game/rpcrepository"
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
	repoBackend := inmemoryrepository.New()
	rpcrepository.RegisterRepoServer(grpcServer, gamerepo.NewServer(repoBackend))

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
