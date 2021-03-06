package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid"
)

// (POST /{game}/play)
func (s *server) Play(ctx echo.Context, g string, params tictactoeapi.PlayParams) error {
	gameID, err := s.verifyID(ctx, g)
	if err != nil {
		return err
	}

	authorizedGameID, mark, err := s.tokenValidator.Validate(params.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if gameID != authorizedGameID {
		return echo.NewHTTPError(http.StatusUnauthorized, "not authorized for this game")
	}

	err = s.turn.TakeTurn(
		ctx.Request().Context(),
		gameID,
		player.Mark(mark),
		grid.Position{
			X: int(params.Position.I),
			Y: int(params.Position.J),
		},
	)
	if err != nil {
		return err
	}

	return ctx.JSON(200, "ok")
}
