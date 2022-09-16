package server_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
)

func TestPlay(t *testing.T) {
	env := newEnv(t)
	client := env.Client

	player1, player2 := createGame(t, env)

	playRes, err := client.PlayWithResponse(context.Background(), player1.GameID, &tictactoeapi.PlayParams{
		Token: player1.Token,
		I:     0,
		J:     0,
	})
	checkResponse(t, playRes, http.StatusOK, err)

	playRes, err = client.PlayWithResponse(context.Background(), player2.GameID, &tictactoeapi.PlayParams{
		Token: player2.Token,
		I:     0,
		J:     1,
	})
	checkResponse(t, playRes, http.StatusOK, err)

	// Get grid content
	gridRes, err := client.GameGridWithResponse(context.Background(), string(player1.GameID))
	checkResponse(t, gridRes, http.StatusOK, err)
	gridOut := *gridRes.JSON200

	// Check expected positions were set
	if gridOut.Grid[0][0] != player1.Mark {
		t.Errorf("Expected grid to be %q, got %q", player1.Mark, gridOut.Grid[0][0])
	}
	if gridOut.Grid[0][1] != player2.Mark {
		t.Errorf("Expected grid to be %q, got %q", player2.Mark, gridOut.Grid[0][1])
	}
}

func TestPlayOutOfTurn(t *testing.T) {
	env := newEnv(t)
	client := env.Client

	_, player2 := createGame(t, env)

	playRes, err := client.PlayWithResponse(context.Background(), player2.GameID, &tictactoeapi.PlayParams{
		Token: player2.Token,
		I:     0,
		J:     0,
	})
	// TODO: This should return a different error
	checkResponse(t, playRes, http.StatusInternalServerError, err)
}

func TestPlayWrongGame(t *testing.T) {
	env := newEnv(t)
	client := env.Client

	player1, _ := createGame(t, env)
	player1a, _ := createGame(t, env)

	playRes, err := client.PlayWithResponse(context.Background(), player1a.GameID, &tictactoeapi.PlayParams{
		Token: player1.Token,
		I:     0,
		J:     0,
	})
	checkResponse(t, playRes, http.StatusUnauthorized, err)
}

func TestPlayInvalidToken(t *testing.T) {
	env := newEnv(t)
	client := env.Client

	player1, _ := createGame(t, env)

	playRes, err := client.PlayWithResponse(context.Background(), player1.GameID, &tictactoeapi.PlayParams{
		Token: "invalid",
		I:     0,
		J:     0,
	})
	checkResponse(t, playRes, http.StatusUnauthorized, err)
}
