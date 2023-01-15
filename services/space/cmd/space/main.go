package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/env"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/opentelemetry"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui/rpcserver"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	space "github.com/theothertomelliott/tic-tac-toverengineered/services/space/internal"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg/rpcspace"
)

func main() {
	version.Println()

	port, err := env.GetInt("PORT", 8080)
	if err != nil {
		log.Fatalf("could not get port number:  %v", err)
	}

	x, y, err := getPosition()
	if err != nil {
		log.Fatalf("loading position from env:  %v", err)
	}

	otelCleanup, err := opentelemetry.Setup(fmt.Sprintf("space-%v", port))
	if err != nil {
		log.Fatalf("could not configure telemetry: %v", err)
	}
	defer otelCleanup()

	spaceBackend, mongoCleanup, err := getMongoSpaceBackend(x, y)
	defer mongoCleanup()
	if err != nil {
		log.Fatal(err)
	}

	rpcServer := rpcserver.New(port)
	rpcspace.RegisterSpaceServer(rpcServer.GRPC(), space.NewServer(spaceBackend))

	log.Printf("gRPC listening on port :%v", port)
	if err := rpcServer.Serve(); err != nil {
		log.Fatal(err)
	}
}

func getPosition() (int, int, error) {
	x, err := strconv.Atoi(os.Getenv("XPOS"))
	if err != nil {
		return 0, 0, err
	}
	y, err := strconv.Atoi(os.Getenv("YPOS"))
	if err != nil {
		return 0, 0, err
	}
	return x, y, nil
}
