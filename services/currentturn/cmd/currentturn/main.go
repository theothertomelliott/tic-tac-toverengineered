package main

import (
	"log"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/defaultmonitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui/rpcserver"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/currentturn/internal/currentturn"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/currentturn/pkg/turn/inmemoryturns"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/currentturn/pkg/turn/rpcturn"
)

func main() {
	version.Println()
	defaultmonitoring.Init("currentturn")
	defer monitoring.Close()

	port := 8080
	grpcuiPort := 8081

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
