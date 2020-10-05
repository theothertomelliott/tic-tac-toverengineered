package main

import (
	"fmt"
	"log"
	"net"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpcui"
	gridserver "github.com/theothertomelliott/tic-tac-toverengineered/grid/internal/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid/rpcgrid"
	space "github.com/theothertomelliott/tic-tac-toverengineered/space/pkg"
	"github.com/theothertomelliott/tic-tac-toverengineered/space/pkg/rpcspace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 8080
	grpcuiPort := 8081
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

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

	rpcgrid.RegisterGridServer(grpcServer, gridserver.NewServer(gridBackend))

	// we need the reflection service, to power the UI
	reflection.Register(grpcServer)
	log.Printf("gRPC listening on port :%v", port)

	var done = make(chan struct{})
	go func() {
		err := grpcServer.Serve(lis)
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
