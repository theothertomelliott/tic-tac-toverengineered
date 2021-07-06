package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/checker/pkg/win"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/currentturn/pkg/turn"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/matchmaker"
)

var _ tictactoeapi.ServerInterface = &server{}

func New(
	repo game.Repository,
	matchmaker matchmaker.MatchMakerClient,
	turn turn.Controller,
	grid grid.Grid,
	checker win.Checker,
	tokenValidator matchmaker.TokenValidator,
) tictactoeapi.ServerInterface {
	return &server{
		repo:           repo,
		matchmaker:     matchmaker,
		turn:           turn,
		grid:           grid,
		checker:        checker,
		tokenValidator: tokenValidator,
	}
}

type server struct {
	repo           game.Repository
	matchmaker     matchmaker.MatchMakerClient
	turn           turn.Controller
	grid           grid.Grid
	checker        win.Checker
	tokenValidator matchmaker.TokenValidator
}

// (GET /)
func (s *server) Index(ctx echo.Context, params tictactoeapi.IndexParams) error {
	var max int64 = 10
	if params.Max != nil {
		max = *params.Max
	}

	var offset int64 = 0
	if params.Offset != nil {
		offset = *params.Offset
	}

	games, err := s.repo.List(ctx.Request().Context(), max, offset)
	if err != nil {
		return err
	}
	return ctx.JSON(200, games)
}

func (s *server) verifyID(ctx echo.Context, id string) (game.ID, error) {
	gameID := game.ID(id)

	// Verify this game exists
	exists, err := s.repo.Exists(ctx.Request().Context(), gameID)
	if err != nil {
		return game.ID(""), echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if !exists {
		return game.ID(""), echo.NewHTTPError(
			http.StatusNotFound,
			fmt.Sprintf("game not found: %v", gameID),
		)
	}

	return gameID, nil
}
