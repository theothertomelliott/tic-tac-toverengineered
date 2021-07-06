package server_test

import (
	"context"
	"net/http"
	"testing"
)

func TestWinnerGameInProgress(t *testing.T) {
	env := newEnv(t)
	client := env.Client

	player1, _ := createGame(t, env)

	winnerRes, err := client.WinnerWithResponse(context.Background(), player1.GameID)
	checkResponse(t, winnerRes, http.StatusOK, err)
	if winnerRes.JSON200.Winner != nil {
		t.Errorf("Expected winner to be nil, got %v", *winnerRes.JSON200.Winner)
	}
	if winnerRes.JSON200.Draw != nil && *winnerRes.JSON200.Draw {
		t.Errorf("Expected no draw in ongoing game")
	}
}

func TestWinnerInvalidGame(t *testing.T) {
	env := newEnv(t)
	client := env.Client

	// Create a game
	_, _ = createGame(t, env)

	currentRes, err := client.WinnerWithResponse(context.Background(), "invalid")
	checkResponse(t, currentRes, http.StatusNotFound, err)
}
