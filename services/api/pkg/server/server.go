package server

import (
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
) tictactoeapi.ServerInterface {
	return &server{
		repo:       repo,
		matchmaker: matchmaker,
		turn:       turn,
		grid:       grid,
		checker:    checker,
	}
}

type server struct {
	repo       game.Repository
	matchmaker matchmaker.MatchMakerClient
	turn       turn.Controller
	grid       grid.Grid
	checker    win.Checker
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

// (GET /match)
func (s *server) MatchStatus(ctx echo.Context, params tictactoeapi.MatchStatusParams) error {
	res, err := s.matchmaker.Check(ctx.Request().Context(), &matchmaker.CheckRequest{
		RequestId: params.RequestID,
	})
	if err != nil {
		return err
	}
	if res.Match == nil {
		ctx.JSON(102, tictactoeapi.MatchPending{
			RequestID: params.RequestID,
		})
		return nil
	}
	ctx.JSON(200, tictactoeapi.Match{
		GameID: res.Match.GameId,
		Mark:   res.Match.Mark,
		Token:  res.Match.Token,
	})
	return nil
}

// (POST /match)
func (s *server) RequestMatch(ctx echo.Context) error {
	res, err := s.matchmaker.Request(ctx.Request().Context(), &matchmaker.RequestRequest{})
	if err != nil {
		return err
	}
	pending := tictactoeapi.MatchPending{
		RequestID: res.RequestId,
	}
	ctx.JSON(202, pending)
	return nil
}

// (GET /{game}/grid)
func (s *server) GameGrid(ctx echo.Context, game string) error {
	panic("not implemented") // TODO: Implement
}

// (GET /{game}/player/current)
func (s *server) CurrentPlayer(ctx echo.Context, game string) error {
	panic("not implemented") // TODO: Implement
}

// (GET /{game}/player/winner)
func (s *server) Winner(ctx echo.Context, game string) error {
	panic("not implemented") // TODO: Implement
}
