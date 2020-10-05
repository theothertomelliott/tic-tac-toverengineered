package turn

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
)

type Controller interface {
	TakeTurn(context.Context, game.ID, player.Mark, grid.Position) error
	NextPlayer(context.Context, game.ID) (player.Mark, error)
}
