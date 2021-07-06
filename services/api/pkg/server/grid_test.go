package server_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid"
)

func TestGrid(t *testing.T) {
	env := newEnv(t)
	client := env.Client
	repo := env.Repo
	memoryGrid := env.Grid

	gameID, err := repo.New(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Get grid content
	gridRes, err := client.GameGridWithResponse(context.Background(), string(gameID))
	checkResponse(t, gridRes, http.StatusOK, err)
	gridOut := *gridRes.JSON200

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

	// Get grid content
	gridRes, err = client.GameGridWithResponse(context.Background(), string(gameID))
	checkResponse(t, gridRes, http.StatusOK, err)
	gridOut = *gridRes.JSON200

	// Check expected position was set
	if gridOut.Grid[0][0] != "O" {
		t.Errorf("Expected grid to be %q, got %q", "O", gridOut.Grid[0][0])
	}

}
