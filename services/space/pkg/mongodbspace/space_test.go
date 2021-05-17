package mongodbspace_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/mongodbtest"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg/mongodbspace"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func TestMain(m *testing.M) {
	// Parse flags to get testing.Short
	flag.Parse()
	if testing.Short() {
		fmt.Println("Will not run Mongo tests (to enable, remove the short flag)")
		os.Exit(0)
	}

	var (
		cleanup func() error
		err     error
	)

	client, cleanup, err = mongodbtest.DockerClient()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func TestWriteSpace(t *testing.T) {
	collection := client.Database("tictactoe").Collection("spaces")

	s, err := mongodbspace.New(context.Background(), collection, 0, 0)
	fatalIf(t, err)

	gameID := game.ID("mygame")

	m, err := s.Mark(context.Background(), gameID)
	fatalIf(t, err)
	if m != nil {
		t.Errorf("expected nil, got %v", *m)
	}

	err = s.SetMark(context.Background(), gameID, player.O)
	fatalIf(t, err)

	m, err = s.Mark(context.Background(), gameID)
	fatalIf(t, err)
	if m == nil || *m != player.O {
		t.Errorf("expected O, got %v", m)
	}

	err = s.SetMark(context.Background(), gameID, player.X)
	fatalIf(t, err)

	m, err = s.Mark(context.Background(), gameID)
	fatalIf(t, err)
	if m == nil || *m != player.X {
		t.Errorf("expected X, got %v", m)
	}
}

func fatalIf(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
