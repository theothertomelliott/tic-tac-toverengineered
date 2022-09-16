package server_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
)

func TestCurrentPlayer(t *testing.T) {
	env := newEnv(t)
	client := env.Client

	player1, player2 := createGame(t, env)

	currentRes, err := client.CurrentPlayerWithResponse(context.Background(), player1.GameID)
	checkResponse(t, currentRes, http.StatusOK, err)
	if *currentRes.JSON200 != player1.Mark {
		t.Errorf("Expected current player to be %s, got %s", player1.Mark, *currentRes.JSON200)
	}

	playRes, err := client.PlayWithResponse(context.Background(), player1.GameID, &tictactoeapi.PlayParams{
		Token: player1.Token,
		I:     0,
		J:     0,
	})
	checkResponse(t, playRes, http.StatusOK, err)

	currentRes, err = client.CurrentPlayerWithResponse(context.Background(), player1.GameID)
	checkResponse(t, currentRes, http.StatusOK, err)
	if *currentRes.JSON200 != player2.Mark {
		t.Errorf("Expected current player to be %s, got %s", player1.Mark, *currentRes.JSON200)
	}
}

func TestCurrentPlayerInvalidGame(t *testing.T) {
	env := newEnv(t)
	client := env.Client

	// Create a game
	_, _ = createGame(t, env)

	currentRes, err := client.CurrentPlayerWithResponse(context.Background(), "invalid")
	checkResponse(t, currentRes, http.StatusNotFound, err)
}
