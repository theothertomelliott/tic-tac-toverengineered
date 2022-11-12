package main

import (
	"log"
	"os"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/env"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/opentelemetry"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui/rpcserver"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	checkerserver "github.com/theothertomelliott/tic-tac-toverengineered/services/checker/internal/checker"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/checker/pkg/win/gridchecker"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/checker/pkg/win/rpcchecker"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid/rpcgrid"
)

func getGridServerTarget() string {
	if serverTarget := os.Getenv("GRID_SERVER_TARGET"); serverTarget != "" {
		return serverTarget
	}
	return "localhost:8086"
}

func main() {
	version.Println()
	cleanup, err := opentelemetry.Setup("checker")
	if err != nil {
		log.Fatalf("could not configure telemetry: %v", err)
	}
	defer cleanup()

	port := env.MustGetInt("PORT", 8080)
	grpcuiPort := env.MustGetInt("GRPCUI_PORT", 8081)

	g, err := rpcgrid.ConnectGrid(getGridServerTarget())
	if err != nil {
		log.Fatalf("could not connect to grid server: %v", err)
	}
	checkerBackend := gridchecker.New(g)

	rpcServer := rpcserver.New(port)
	rpcchecker.RegisterCheckerServer(rpcServer.GRPC(), checkerserver.NewServer(checkerBackend))

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
