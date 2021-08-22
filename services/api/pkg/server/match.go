package server

import (
	"net/http"

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
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if res.Match == nil {
		ctx.JSON(202, tictactoeapi.MatchPending{
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

// (POST /match/pair)
func (s *server) RequestMatchPair(ctx echo.Context) error {
	res, err := s.matchmaker.RequestPair(ctx.Request().Context(), &matchmaker.RequestPairRequest{})
	if err != nil {
		return err
	}
	out := tictactoeapi.MatchPair{
		X: tictactoeapi.Match{
			GameID: res.X.GameId,
			Mark:   res.X.Mark,
			Token:  res.X.Token,
		},
		O: tictactoeapi.Match{
			GameID: res.O.GameId,
			Mark:   res.O.Mark,
			Token:  res.O.Token,
		},
	}
	ctx.JSON(200, out)
	return nil
}
