package server

import (
	"github.com/labstack/echo/v4"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/matchmaker"
)

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
