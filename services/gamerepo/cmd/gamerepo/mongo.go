package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game/mongodbrepository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

func getMongoGameRepositoryBackend() (game.Repository, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	opts := options.Client()
	opts.Monitor = otelmongo.NewMonitor()
	opts.ApplyURI(os.Getenv("MONGO_CONN"))
	client, err := mongo.Connect(
		ctx,
		opts,
	)
	if err != nil {
		return nil, cancel, fmt.Errorf("connecting to mongo:  %v", err)
	}

	cleanup := func() {
		cancel()
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("disconnecting from mongo:  %v", err)
		}
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, cancel, fmt.Errorf("checking mongo connection:  %v", err)
	}

	repoBackend, err := mongodbrepository.New(
		context.Background(),
		client.Database("tictactoe").Collection("games"),
	)
	if err != nil {
		return nil, cleanup, err
	}
	return repoBackend, cleanup, err
}
