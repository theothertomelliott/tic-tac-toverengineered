package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/env"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/opentelemetry"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui/rpcserver"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	gridserver "github.com/theothertomelliott/tic-tac-toverengineered/services/grid/internal/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid/rpcgrid"
	space "github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg/rpcspace"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	version.Println()
	cleanup, err := opentelemetry.Setup("grid")
	if err != nil {
		log.Fatalf("could not configure telemetry: %v", err)
	}
	defer cleanup()

	port := env.MustGetInt("PORT", 8080)
	grpcuiPort := env.MustGetInt("GRPCUI_PORT", 8081)

	var (
		spaces       [][]space.Space
		healthChecks [][]grpc_health_v1.HealthClient
	)
	for i := 0; i < 3; i++ {
		var (
			row       []space.Space
			healthRow []grpc_health_v1.HealthClient
		)
		for j := 0; j < 3; j++ {
			connStr := env.Get(fmt.Sprintf("SPACE-%d-%d", i, j), fmt.Sprintf("localhost:80%v%v", i, j))
			c, err := rpcspace.ConnectSpace(connStr)
			if err != nil {
				log.Fatalf("space (%d,%d): %v", i, j, err)
			}
			row = append(row, c)
			healthRow = append(healthRow, c.Health())
		}
		spaces = append(spaces, row)
		healthChecks = append(healthChecks, healthRow)
	}
	gridBackend, _ := grid.New(spaces)

	hs := health.NewServer()
	rpcServer := rpcserver.NewWithHealthServer(port, hs)
	rpcgrid.RegisterGridServer(rpcServer.GRPC(), gridserver.NewServer(gridBackend))

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

	go checkHealth(context.Background(), healthChecks, hs)

	<-done
}

func checkHealth(ctx context.Context, spaces [][]grpc_health_v1.HealthClient, hs *health.Server) {
	for {
		var serving = true
		for _, row := range spaces {
			for _, space := range row {
				res, err := space.Check(ctx, &grpc_health_v1.HealthCheckRequest{})
				if err != nil {
					serving = false
					continue
				}
				if res.Status != grpc_health_v1.HealthCheckResponse_SERVING {
					serving = false
				}
			}
		}
		if serving {
			hs.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
		} else {
			hs.SetServingStatus("", grpc_health_v1.HealthCheckResponse_NOT_SERVING)
		}
		time.Sleep(time.Second)
	}
}
