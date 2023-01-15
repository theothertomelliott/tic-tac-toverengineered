package main

import (
	"log"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/env"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/opentelemetry"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui/rpcserver"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/internal/gamerepo"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game/rpcrepository"
)

func main() {
	version.Println()
	cleanup, err := opentelemetry.Setup("gamerepo")
	if err != nil {
		log.Fatalf("could not configure telemetry: %v", err)
	}
	defer cleanup()

	port := env.MustGetInt("PORT", 8080)
	grpcuiPort := env.MustGetInt("GRPCUI_PORT", 8081)

	repoBackend, backendCleanup, err := getMongoGameRepositoryBackend()
	defer backendCleanup()
	if err != nil {
		log.Fatalf("could not create backend: %v", err)
	}

	rpcServer := rpcserver.New(port)
	rpcrepository.RegisterRepoServer(rpcServer.GRPC(), gamerepo.NewServer(repoBackend))

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
