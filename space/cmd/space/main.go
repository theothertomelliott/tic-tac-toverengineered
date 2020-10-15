package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui/rpcserver"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	space "github.com/theothertomelliott/tic-tac-toverengineered/space/internal"
	"github.com/theothertomelliott/tic-tac-toverengineered/space/pkg/rpcspace"
	"github.com/theothertomelliott/tic-tac-toverengineered/space/pkg/spaceinmemory"
)

func getPort() (int, error) {
	if serverTarget := os.Getenv("PORT"); serverTarget != "" {
		return strconv.Atoi(serverTarget)
	}
	return 8080, nil
}

func main() {
	version.Println()

	port, err := getPort()
	if err != nil {
		log.Fatalf("could not get port number:  %v", err)
	}

	spaceBackend := spaceinmemory.New()

	rpcServer := rpcserver.New(port)
	rpcspace.RegisterSpaceServer(rpcServer.GRPC(), space.NewServer(spaceBackend))

	closeMonitoring := monitoring.Init(fmt.Sprintf("space-%v", port))
	defer closeMonitoring()

	log.Printf("gRPC listening on port :%v", port)
	if err := rpcServer.Serve(); err != nil {
		log.Fatal(err)
	}
}
