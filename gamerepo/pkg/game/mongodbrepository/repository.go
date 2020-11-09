package mongodbrepository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// New creates a new game.Repository with
// games recorded in memory.
func New(ctx context.Context, collection *mongo.Collection) (game.Repository, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Add unique index on Game and Position
	name, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "game", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("Created index: ", name)

	return &repository{
		collection: collection,
	}, nil
}

type repository struct {
	collection *mongo.Collection
}

// New creates a new game and creates a unique ID
func (r *repository) New(ctx context.Context) (game.ID, error) {
	u := uuid.New()
	gameID := game.ID(u.String())
	_, err := r.collection.InsertOne(ctx, bson.M{
		"game": gameID,
	})
	if err != nil {
		return game.ID(""), err
	}
	return gameID, nil
}

// List obtains game IDs, ordered by creation date.
// Pagination is managed through the max and offset params.
func (r *repository) List(ctx context.Context, max int64, offset int64) ([]game.ID, error) {
	cursor, err := r.collection.Find(
		ctx,
		struct{}{},
		options.Find().SetLimit(max).SetSkip(offset).SetSort(bson.D{{"game", 1}}),
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var games []game.ID
	for cursor.Next(ctx) {
		var result struct {
			Game game.ID
		}
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		games = append(games, result.Game)
	}
	return games, nil
}

// Exists returns true iff the given game ID was previously created with New
func (r *repository) Exists(ctx context.Context, gameID game.ID) (bool, error) {
	var filter = struct {
		Game game.ID
	}{
		Game: gameID,
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	res := r.collection.FindOne(ctx, filter)
	if res.Err() == mongo.ErrNoDocuments {
		return false, nil
	}
	if res.Err() != nil {
		return false, res.Err()
	}
	return true, nil
}
