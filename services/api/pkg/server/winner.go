package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
)

// (GET /{game}/player/winner)
func (s *server) Winner(ctx echo.Context, g string) error {
	gameID, err := s.verifyID(ctx, g)
	if err != nil {
		return err
	}

	winner, err := s.checker.Winner(ctx.Request().Context(), gameID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	result := tictactoeapi.Winner{
		Draw: &winner.IsDraw,
	}
	if winner.Winner != nil {
		w := fmt.Sprint(*winner.Winner)
		result.Winner = &w
	}

	return ctx.JSON(200, result)
}
