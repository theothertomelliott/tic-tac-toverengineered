package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui/rpcserver"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	space "github.com/theothertomelliott/tic-tac-toverengineered/space/internal"
	"github.com/theothertomelliott/tic-tac-toverengineered/space/pkg/mongodbspace"
	"github.com/theothertomelliott/tic-tac-toverengineered/space/pkg/rpcspace"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func getPort() (int, error) {
	if serverTarget := os.Getenv("PORT"); serverTarget != "" {
		return strconv.Atoi(serverTarget)
	}
	return 8080, nil
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

func main() {
	version.Println()

	port, err := getPort()
	if err != nil {
		log.Fatalf("could not get port number:  %v", err)
	}

	x, y, err := getPosition()
	if err != nil {
		log.Fatalf("loading position from env:  %v", err)
	}

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

	closeMonitoring := monitoring.Init(fmt.Sprintf("space-%v", port))
	defer closeMonitoring()

	log.Printf("gRPC listening on port :%v", port)
	if err := rpcServer.Serve(); err != nil {
		log.Fatal(err)
	}
}
