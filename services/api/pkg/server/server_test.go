package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/server"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/checker/pkg/win/gridchecker"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/currentturn/pkg/turn"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/currentturn/pkg/turn/inmemoryturns"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game/inmemoryrepository"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/matchmaker/inmemorymatchmaker"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/matchmaker/unsignedtokens"
)

type env struct {
	Client tictactoeapi.ClientWithResponsesInterface
	Repo   game.Repository
	Grid   grid.Grid
	Turns  turn.Controller
}

// newEnv creates a testing environment with a client connected to an in-memory server
func newEnv(t *testing.T) *env {
	gamerepo := inmemoryrepository.New()
	memoryGrid := grid.NewInMemory()
	current := inmemoryturns.NewCurrentTurn()
	checker := gridchecker.New(memoryGrid)
	turns := inmemoryturns.New(current, memoryGrid, checker)
	apiServer := server.New(
		gamerepo,
		inmemorymatchmaker.New(gamerepo),
		turns,
		memoryGrid,
		checker,
		&unsignedtokens.UnsignedTokens{})

	e := echo.New()
	tictactoeapi.RegisterHandlers(e, apiServer)

	client, err := tictactoeapi.NewClientWithResponses("", tictactoeapi.WithHTTPClient(&echoDoer{
		echo: e,
	}))
	if err != nil {
		t.Fatal(err)
	}
	return &env{
		Client: client,
		Repo:   gamerepo,
		Grid:   memoryGrid,
		Turns:  turns,
	}
}

type echoDoer struct {
	echo *echo.Echo
}

func (e *echoDoer) Do(req *http.Request) (*http.Response, error) {
	var res = httptest.NewRecorder()
	e.echo.ServeHTTP(res, req)
	return res.Result(), nil
}

func checkResponse(t *testing.T, res resWithStatusCode, expectedCode int, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode() != expectedCode {
		t.Errorf("Expected status code %v, got %v", expectedCode, res.StatusCode())
	}
}

type resWithStatusCode interface {
	StatusCode() int
}
