package server_test

import (
	"context"
	"testing"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/server"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game/inmemoryrepository"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/matchmaker/inmemorymatchmaker"
)

func TestGrid(t *testing.T) {
	gamerepo := inmemoryrepository.New()
	memoryGrid := grid.NewInMemory()
	apiServer := server.New(
		gamerepo,
		inmemorymatchmaker.New(gamerepo),
		nil,
		memoryGrid,
		nil,
	)

	gameID, err := gamerepo.New(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	var gridOut tictactoeapi.Grid
	request(
		t,
		gridReq(string(gameID)),
		gridCall(t, apiServer, string(gameID)),
		200,
		&gridOut,
	)

	// check gridOut is all empty strings
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if gridOut.Grid[i][j] != "" {
				t.Errorf("Expected grid to be empty, got %q", gridOut.Grid[i][j])
			}
		}
	}

	// Set a position
	memoryGrid.SetMark(context.Background(), gameID, grid.Position{X: 0, Y: 0}, player.O)
	request(
		t,
		gridReq(string(gameID)),
		gridCall(t, apiServer, string(gameID)),
		200,
		&gridOut,
	)
	if gridOut.Grid[0][0] != "O" {
		t.Errorf("Expected grid to be %q, got %q", "O", gridOut.Grid[0][0])
	}

}
