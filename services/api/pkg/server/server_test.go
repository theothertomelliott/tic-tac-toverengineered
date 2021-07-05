package server_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
)

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

func gridReq(gameID string) *http.Request {
	return httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%v/grid", gameID), nil)
}

func gridCall(t *testing.T, apiServer tictactoeapi.ServerInterface, gameID string) func(ctx echo.Context) {
	return func(ctx echo.Context) {
		if err := apiServer.GameGrid(ctx, gameID); err != nil {
			t.Error(err)
		}
	}
}
