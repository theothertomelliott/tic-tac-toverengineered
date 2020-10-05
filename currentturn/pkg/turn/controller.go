package turn

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
)

type Controller interface {
	TakeTurn(context.Context, game.ID, player.Mark, grid.Position) error
	NextPlayer(context.Context, game.ID) (player.Mark, error)
}
