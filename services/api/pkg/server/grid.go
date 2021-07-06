package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
)

// (GET /{game}/grid)
func (s *server) GameGrid(ctx echo.Context, game string) error {
	gameID, err := s.verifyID(ctx, game)
	if err != nil {
		return err
	}

	out, err := s.grid.State(ctx.Request().Context(), gameID)
	if err != nil {
		fmt.Println("error getting grid state: ", err)
		return err
	}
	var gridOut tictactoeapi.Grid
	for _, row := range out {
		gridOut.Grid = append(gridOut.Grid, make([]string, len(row)))
		for i, cell := range row {
			if cell == nil {
				gridOut.Grid[len(gridOut.Grid)-1][i] = ""
			} else {
				gridOut.Grid[len(gridOut.Grid)-1][i] = cell.String()
			}
		}
	}

	return ctx.JSON(200, gridOut)
}
