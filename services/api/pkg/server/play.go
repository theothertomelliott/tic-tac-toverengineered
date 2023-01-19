package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/logctx"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid"
)

// (POST /{game}/play)
func (s *server) Play(ectx echo.Context, g string, params tictactoeapi.PlayParams) error {
	ctx := ectx.Request().Context()
	gameID, err := s.verifyID(ectx, g)
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
		ctx,
		gameID,
		player.Mark(mark),
		grid.Position{
			X: int(params.I),
			Y: int(params.J),
		},
	)
	if err != nil {
		logctx.Printf(ctx, "error taking turn: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("turn could not be made: %v", err))
	}

	return ectx.JSON(200, "ok")
}
