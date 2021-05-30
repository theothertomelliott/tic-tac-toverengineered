package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/server"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game/inmemoryrepository"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/matchmaker/inmemorymatchmaker"
)

func TestMatching(t *testing.T) {
	gamerepo := inmemoryrepository.New()
	apiServer := server.New(gamerepo, inmemorymatchmaker.New(gamerepo), nil, nil, nil)

	pending1 := &tictactoeapi.MatchPending{}
	request(
		t,
		matchRequest(),
		matchCall(t, apiServer),
		202,
		pending1,
	)
	pending2 := &tictactoeapi.MatchPending{}
	request(
		t,
		matchRequest(),
		matchCall(t, apiServer),
		202,
		pending2,
	)

	if pending1.RequestID == pending2.RequestID {
		t.Errorf("Request IDs should not match")
	}

	match1 := &tictactoeapi.Match{}
	request(
		t,
		matchStatus(),
		matchStatusCall(t, apiServer, pending1.RequestID),
		200,
		match1,
	)

	match2 := &tictactoeapi.Match{}
	request(
		t,
		matchStatus(),
		matchStatusCall(t, apiServer, pending2.RequestID),
		200,
		match2,
	)

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
	gameList := []string{}
	request(
		t,
		index(),
		indexCall(t, apiServer),
		200,
		&gameList,
	)

	if len(gameList) != 1 {
		t.Errorf("Expected one game, got %v", len(gameList))
	} else if gameList[0] != match1.GameID {
		t.Errorf("Expected listed game ID to be %q, got %q", match1.GameID, gameList[0])
	}

}

func request(t *testing.T, req *http.Request, callback func(ctx echo.Context), expectedCode int, out interface{}) {
	e := echo.New()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	callback(c)

	if rec.Code != expectedCode {
		t.Errorf("expected %v for new request, got %v", expectedCode, rec.Code)
	}

	body := rec.Body.Bytes()
	err := json.Unmarshal(body, out)
	if err != nil {
		t.Error(err)
	}
}

func matchRequest() *http.Request {
	return httptest.NewRequest(http.MethodPost, "/match", nil)
}

func matchCall(t *testing.T, apiServer tictactoeapi.ServerInterface) func(ctx echo.Context) {
	return func(ctx echo.Context) {
		if err := apiServer.RequestMatch(ctx); err != nil {
			t.Error(err)
		}
	}
}

func matchStatus() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/match", nil)
}

func matchStatusCall(t *testing.T, apiServer tictactoeapi.ServerInterface, requestID string) func(ctx echo.Context) {
	return func(ctx echo.Context) {
		if err := apiServer.MatchStatus(ctx, tictactoeapi.MatchStatusParams{
			RequestID: requestID,
		}); err != nil {
			t.Error(err)
		}
	}
}

func index() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/", nil)
}

func indexCall(t *testing.T, apiServer tictactoeapi.ServerInterface) func(ctx echo.Context) {
	return func(ctx echo.Context) {
		if err := apiServer.Index(ctx, tictactoeapi.IndexParams{}); err != nil {
			t.Error(err)
		}
	}
}
