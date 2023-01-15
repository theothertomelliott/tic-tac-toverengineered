package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	space "github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg/mongodbspace"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

func getMongoSpaceBackend(x, y int) (space.Space, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	opts := options.Client()
	opts.Monitor = otelmongo.NewMonitor()
	opts.ApplyURI(os.Getenv("MONGO_CONN"))
	client, err := mongo.Connect(
		ctx,
		opts,
	)
	if err != nil {
		return nil, func() {
			cancel()
		}, fmt.Errorf("connecting to mongo:  %v", err)
	}

	cleanup := func() {
		cancel()
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("disconnecting from mongo:  %v", err)
		}
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, func() {
			cancel()
		}, fmt.Errorf("checking mongo connection:  %v", err)
	}

	space, err := mongodbspace.New(
		context.Background(),
		client.Database("tictactoe").Collection("spaces"),
		x,
		y,
	)
	if err != nil {
		return nil, cleanup, err
	}

	return space, cleanup, nil
}
