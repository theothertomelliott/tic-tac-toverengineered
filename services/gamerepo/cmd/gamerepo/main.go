package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/env"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/opentelemetry"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui/rpcserver"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/internal/gamerepo"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game/mongodbrepository"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game/rpcrepository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

func main() {
	version.Println()
	cleanup, err := opentelemetry.Setup("gamerepo")
	if err != nil {
		log.Fatalf("could not configure telemetry: %v", err)
	}
	defer cleanup()

	port := env.MustGetInt("PORT", 8080)
	grpcuiPort := env.MustGetInt("GRPCUI_PORT", 8081)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client()
	opts.Monitor = otelmongo.NewMonitor()
	opts.ApplyURI(os.Getenv("MONGO_CONN"))
	client, err := mongo.Connect(
		ctx,
		opts,
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
