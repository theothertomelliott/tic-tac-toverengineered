package main

import (
	"log"
	"os"

	"github.com/theothertomelliott/tic-tac-toverengineered/checker/pkg/win/rpcchecker"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/defaultmonitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui/rpcserver"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	"github.com/theothertomelliott/tic-tac-toverengineered/currentturn/pkg/turn/inmemoryturns"
	"github.com/theothertomelliott/tic-tac-toverengineered/currentturn/pkg/turn/rpcturn"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid/rpcgrid"
	"github.com/theothertomelliott/tic-tac-toverengineered/turncontroller/internal/turncontroller"
)

func getCurrentTurnServerTarget() string {
	if serverTarget := os.Getenv("CURRENT_TURN_SERVER_TARGET"); serverTarget != "" {
		return serverTarget
	}
	return "localhost:8084"
}

func getGridServerTarget() string {
	if serverTarget := os.Getenv("GRID_SERVER_TARGET"); serverTarget != "" {
		return serverTarget
	}
	return "localhost:8086"
}

func getCheckerServerTarget() string {
	if serverTarget := os.Getenv("CHECKER_SERVER_TARGET"); serverTarget != "" {
		return serverTarget
	}
	return "localhost:8088"
}

func main() {
	version.Println()

	port := 8080
	grpcuiPort := 8081

	g, err := rpcgrid.ConnectGrid(getGridServerTarget())
	if err != nil {
		log.Fatalf("could not connect to grid server: %v", err)
	}
	checker, err := rpcchecker.ConnectChecker(getCheckerServerTarget())
	if err != nil {
		log.Fatalf("could not connect to win checker server: %v", err)
	}
	ct, err := rpcturn.ConnectCurrent(getCurrentTurnServerTarget())
	if err != nil {
		log.Fatalf("could not connect to current turn server: %v", err)
	}
	controllerBackend := inmemoryturns.New(ct, g, checker)

	rpcServer := rpcserver.New(port)
	rpcturn.RegisterControllerServer(rpcServer.GRPC(), turncontroller.NewServer(controllerBackend))

	defaultmonitoring.Init("turncontroller")
	defer monitoring.Close()

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
