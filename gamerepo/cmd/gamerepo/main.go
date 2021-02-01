package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/defaultmonitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui/rpcserver"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/internal/gamerepo"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game/mongodbrepository"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game/rpcrepository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	version.Println()
	defaultmonitoring.Init("gamerepo")
	defer monitoring.Close()

	port := 8080
	grpcuiPort := 8081

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

	repoBackend, err := mongodbrepository.New(
		context.Background(),
		client.Database("tictactoe").Collection("games"),
	)

	rpcServer := rpcserver.New(port)
	rpcrepository.RegisterRepoServer(rpcServer.GRPC(), gamerepo.NewServer(repoBackend))

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
