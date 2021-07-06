package server

import (
	"github.com/labstack/echo/v4"
)

// (GET /{game}/player/current)
func (s *server) CurrentPlayer(ctx echo.Context, g string) error {
	gameID, err := s.verifyID(ctx, g)
	if err != nil {
		return err
	}

	mark, err := s.turn.NextPlayer(ctx.Request().Context(), gameID)
	if err != nil {
		return err
	}

	return ctx.JSON(200, mark)
}
