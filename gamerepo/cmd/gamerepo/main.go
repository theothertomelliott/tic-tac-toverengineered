package main

import (
	"log"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui/rpcserver"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/internal/gamerepo"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game/inmemoryrepository"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game/rpcrepository"
)

func main() {
	version.Println()

	port := 8080
	grpcuiPort := 8081

	repoBackend := inmemoryrepository.New()

	rpcServer := rpcserver.New(port)
	rpcrepository.RegisterRepoServer(rpcServer.GRPC(), gamerepo.NewServer(repoBackend))

	closeMonitoring := monitoring.Init("gamerepo")
	defer closeMonitoring()

	log.Printf("gRPC listening on port :%v", port)
	var done = make(chan struct{})
	go func() {
		err := rpcServer.Serve()
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
