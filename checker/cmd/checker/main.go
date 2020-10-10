package main

import (
	"log"
	"os"

	checkerserver "github.com/theothertomelliott/tic-tac-toverengineered/checker/internal/checker"
	"github.com/theothertomelliott/tic-tac-toverengineered/checker/pkg/win/gridchecker"
	"github.com/theothertomelliott/tic-tac-toverengineered/checker/pkg/win/rpcchecker"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui/rpcserver"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid/rpcgrid"
)

func getGridServerTarget() string {
	if serverTarget := os.Getenv("GRID_SERVER_TARGET"); serverTarget != "" {
		return serverTarget
	}
	return "localhost:8086"
}

func main() {
	version.Println()

	port := 8080
	grpcuiPort := 8081

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
