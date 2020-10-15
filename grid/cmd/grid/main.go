package main

import (
	"fmt"
	"log"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui/rpcserver"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	gridserver "github.com/theothertomelliott/tic-tac-toverengineered/grid/internal/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid/rpcgrid"
	space "github.com/theothertomelliott/tic-tac-toverengineered/space/pkg"
	"github.com/theothertomelliott/tic-tac-toverengineered/space/pkg/rpcspace"
)

func main() {
	version.Println()

	port := 8080
	grpcuiPort := 8081

	var spaces [][]space.Space
	for i := 0; i < 3; i++ {
		var row []space.Space
		for j := 0; j < 3; j++ {
			c, err := rpcspace.ConnectSpace(fmt.Sprintf("space-%d-%d:80", i, j))
			if err != nil {
				log.Fatalf("space (%d,%d): %v", i, j, err)
			}
			row = append(row, c)
		}
		spaces = append(spaces, row)
	}
	gridBackend, _ := grid.New(spaces)

	rpcServer := rpcserver.New(port)
	rpcgrid.RegisterGridServer(rpcServer.GRPC(), gridserver.NewServer(gridBackend))

	closeMonitoring := monitoring.Init("grid")
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
