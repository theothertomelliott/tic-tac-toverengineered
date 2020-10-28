package mongodbspace

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
	space "github.com/theothertomelliott/tic-tac-toverengineered/space/pkg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ space.Space = &Space{}

// New creates a Space backed by MongoDB.
func New(
	ctx context.Context,
	collection *mongo.Collection,
	x, y int,
) (space.Space, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Add unique index on Game and Position
	name, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "game", Value: 1},
			{Key: "position", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("Created index: ", name)

	return &Space{
		collection: collection,
		pos: position{
			X: x,
			Y: y,
		},
	}, nil
}

type Space struct {
	collection *mongo.Collection
	pos        position
}

func (s *Space) Mark(ctx context.Context, gameID game.ID) (*player.Mark, error) {
	var result struct {
		ID   string `bson:"_id"`
		Mark *player.Mark
	}
	var filter = struct {
		Position position
		Game     game.ID
	}{
		Position: s.pos,
		Game:     gameID,
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	res := s.collection.FindOne(ctx, filter)
	if res.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}
	err := res.Decode(&result)
	return result.Mark, err
}

func (s *Space) SetMark(ctx context.Context, gameID game.ID, m player.Mark) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.collection.UpdateOne(ctx,
		bson.M{
			"position": s.pos,
			"game":     gameID,
		},
		bson.M{
			"$set": bson.M{
				"mark": &m,
			},
		},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return err
	}
	log.Printf("Result %+v", res)
	return nil
}

type position struct {
	X int
	Y int
}
