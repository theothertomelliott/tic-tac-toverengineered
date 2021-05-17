package mongodbrepository_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sort"
	"testing"

	"github.com/google/uuid"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/mongodbtest"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game/mongodbrepository"
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

func TestNewGame(t *testing.T) {
	repoBackend := newRepoBackend(t)

	// Create a game
	gameID, err := repoBackend.New(context.Background())
	fatalIf(t, err)

	// Verify that the game exists
	exists, err := repoBackend.Exists(context.Background(), gameID)
	fatalIf(t, err)
	if !exists {
		t.Error("Expected game exists")
	}

	// Ensure that new games have different IDs
	gameID2, err := repoBackend.New(context.Background())
	fatalIf(t, err)
	if gameID2 == gameID {
		t.Error("second game had same ID as first")
	}

	// Verify that a game that wasn't created doesn't exit
	exists, err = repoBackend.Exists(context.Background(), game.ID("missing"))
	fatalIf(t, err)
	if exists {
		t.Error("Did not expect game to exist")
	}

}

func TestListGames(t *testing.T) {
	repoBackend := newRepoBackend(t)

	var gameIDs []string
	// Create 100 games
	for i := 0; i < 100; i++ {
		gameID, err := repoBackend.New(context.Background())
		fatalIf(t, err)
		gameIDs = append(gameIDs, string(gameID))
	}
	sort.Strings(gameIDs)

	// Read back games 10 at a times
	for i := 0; i < 100; i += 10 {
		games, err := repoBackend.List(context.Background(), 10, int64(i))
		fatalIf(t, err)
		for index, gameID := range games {
			offset := i + index
			expected := gameIDs[offset]
			got := string(gameID)
			if got != expected {
				t.Errorf("%d: expected %q, got %q", offset, expected, got)
			}
		}
	}
}

func newRepoBackend(t *testing.T) game.Repository {
	// Use a uuid to ensure we get a unique collection for each test
	r := uuid.New()
	repoBackend, err := mongodbrepository.New(
		context.Background(),
		client.Database("tictactoe").Collection(fmt.Sprintf("games-%v", r.String())),
	)
	fatalIf(t, err)
	return repoBackend
}

func fatalIf(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
