package main

import (
	"log"
	"os"
	"strconv"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui/rpcserver"
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
	port, err := getPort()
	if err != nil {
		log.Fatalf("could not get port number:  %v", err)
	}

	spaceBackend := spaceinmemory.New()

	rpcServer := rpcserver.New(port)
	rpcspace.RegisterSpaceServer(rpcServer.GRPC(), space.NewServer(spaceBackend))

	log.Printf("gRPC listening on port :%v", port)
	if err := rpcServer.Serve(); err != nil {
		log.Fatal(err)
	}
}
