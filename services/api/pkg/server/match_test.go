package server_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
)

func TestMatching(t *testing.T) {
	env := newEnv(t)
	client := env.Client

	matchRes, err := client.RequestMatchWithResponse(context.Background())
	checkResponse(t, matchRes, http.StatusAccepted, err)
	pending1 := matchRes.JSON202

	// Check match is still pending
	statusRes, err := client.MatchStatusWithResponse(context.Background(), &tictactoeapi.MatchStatusParams{
		RequestID: pending1.RequestID,
	})
	checkResponse(t, statusRes, http.StatusAccepted, err)

	matchRes, err = client.RequestMatchWithResponse(context.Background())
	checkResponse(t, matchRes, http.StatusAccepted, err)
	pending2 := matchRes.JSON202

	if pending1.RequestID == pending2.RequestID {
		t.Errorf("Request IDs should not match")
	}

	statusRes, err = client.MatchStatusWithResponse(context.Background(), &tictactoeapi.MatchStatusParams{
		RequestID: pending1.RequestID,
	})
	checkResponse(t, statusRes, http.StatusOK, err)
	match1 := statusRes.JSON200

	statusRes, err = client.MatchStatusWithResponse(context.Background(), &tictactoeapi.MatchStatusParams{
		RequestID: pending2.RequestID,
	})
	checkResponse(t, statusRes, http.StatusOK, err)
	match2 := statusRes.JSON200

	if match1.GameID != match2.GameID {
		t.Errorf("game ids must match, got %q and %q", match1.GameID, match2.GameID)
	}
	if match1.Mark == match2.Mark {
		t.Errorf("marks must not be the same")
	}
	if match1.Token == match2.Token {
		t.Errorf("tokens must not be the same")
	}

	// Check number of games
	res3, err := client.IndexWithResponse(context.Background(), &tictactoeapi.IndexParams{})
	if err != nil {
		t.Fatal(err)
	}
	if res3.StatusCode() != http.StatusOK {
		t.Errorf("Expected 202, got %d", res3.StatusCode())
	}
	gameList := *res3.JSON200

	if len(gameList) != 1 {
		t.Errorf("Expected one game, got %v", len(gameList))
	} else if gameList[0] != match1.GameID {
		t.Errorf("Expected listed game ID to be %q, got %q", match1.GameID, gameList[0])
	}
}

func createGame(t *testing.T, e *env) (*tictactoeapi.Match, *tictactoeapi.Match) {
	client := e.Client

	matchRes, err := client.RequestMatchWithResponse(context.Background())
	checkResponse(t, matchRes, http.StatusAccepted, err)
	pending1 := matchRes.JSON202

	matchRes, err = client.RequestMatchWithResponse(context.Background())
	checkResponse(t, matchRes, http.StatusAccepted, err)
	pending2 := matchRes.JSON202

	statusRes, err := client.MatchStatusWithResponse(context.Background(), &tictactoeapi.MatchStatusParams{
		RequestID: pending1.RequestID,
	})
	checkResponse(t, statusRes, http.StatusOK, err)
	match1 := statusRes.JSON200

	statusRes, err = client.MatchStatusWithResponse(context.Background(), &tictactoeapi.MatchStatusParams{
		RequestID: pending2.RequestID,
	})
	checkResponse(t, statusRes, http.StatusOK, err)
	match2 := statusRes.JSON200

	return match1, match2
}
