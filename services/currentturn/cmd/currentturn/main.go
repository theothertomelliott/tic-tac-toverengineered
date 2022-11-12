package main

import (
	"log"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/env"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/opentelemetry"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui/rpcserver"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/currentturn/internal/currentturn"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/currentturn/pkg/turn/inmemoryturns"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/currentturn/pkg/turn/rpcturn"
)

func main() {
	version.Println()
	cleanup, err := opentelemetry.Setup("currentturn")
	if err != nil {
		log.Fatalf("could not configure telemetry: %v", err)
	}
	defer cleanup()

	port := env.MustGetInt("PORT", 8080)
	grpcuiPort := env.MustGetInt("GRPCUI_PORT", 8081)

	currentBackend := inmemoryturns.NewCurrentTurn()

	rpcServer := rpcserver.New(port)
	rpcturn.RegisterCurrentServer(rpcServer.GRPC(), currentturn.NewServer(currentBackend))

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
