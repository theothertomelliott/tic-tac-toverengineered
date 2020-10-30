package main

import (
	"context"
	"log"
	"time"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/space/pkg/mongodbspace"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://admin:password@localhost:27017"),
	)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	collection := client.Database("tictactoe").Collection("spaces")

	s, err := mongodbspace.New(context.Background(), collection, 0, 0)
	if err != nil {
		panic(err)
	}

	gameID := game.ID("mygame")

	m, err := s.Mark(context.Background(), gameID)
	if err != nil {
		panic(err)
	}

	log.Println("Mark: ", m)

	err = s.SetMark(context.Background(), gameID, player.O)
	if err != nil {
		panic(err)
	}

	m, err = s.Mark(context.Background(), gameID)
	if err != nil {
		panic(err)
	}

	log.Println("Mark: ", m)

	err = s.SetMark(context.Background(), gameID, player.X)
	if err != nil {
		panic(err)
	}

	m, err = s.Mark(context.Background(), gameID)
	if err != nil {
		panic(err)
	}

	log.Println("Mark: ", m)

	// ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println("Ping succeeded")

	// collection := client.Database("testing").Collection("numbers")

	// ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	// if err != nil {
	// 	panic(err)
	// }
	// id := res.InsertedID
	// log.Printf("Inserted with id %v", id)

	// log.Println("=== Find ===")

	// ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()
	// cur, err := collection.Find(ctx, bson.D{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer cur.Close(ctx)
	// for cur.Next(ctx) {
	// 	var result bson.M
	// 	err := cur.Decode(&result)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Print(result)
	// }
	// if err := cur.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("=== FindOne ===")

	// var result struct {
	// 	ID    string `bson:"_id"`
	// 	Value float64
	// }
	// type nameFilter struct {
	// 	Name string
	// }
	// filter := nameFilter{Name: "pi"}
	// ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// err = collection.FindOne(ctx, filter).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(result)

}
