package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/env"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/opentelemetry"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui/rpcserver"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	space "github.com/theothertomelliott/tic-tac-toverengineered/services/space/internal"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg/mongodbspace"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg/rpcspace"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

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

	cleanup, err := opentelemetry.Setup(fmt.Sprintf("space-%v", port))
	if err != nil {
		log.Fatalf("could not configure telemetry: %v", err)
	}
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(os.Getenv("MONGO_CONN")),
	)
	if err != nil {
		log.Fatalf("connecting to mongo:  %v", err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("disconnecting from mongo:  %v", err)
		}
	}()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("checking mongo connection:  %v", err)
	}

	spaceBackend, err := mongodbspace.New(
		context.Background(),
		client.Database("tictactoe").Collection("spaces"),
		x,
		y,
	)

	rpcServer := rpcserver.New(port)
	rpcspace.RegisterSpaceServer(rpcServer.GRPC(), space.NewServer(spaceBackend))

	log.Printf("gRPC listening on port :%v", port)
	if err := rpcServer.Serve(); err != nil {
		log.Fatal(err)
	}
}
