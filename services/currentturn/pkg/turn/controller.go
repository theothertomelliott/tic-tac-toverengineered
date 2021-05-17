package turn

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid"
)

type Controller interface {
	TakeTurn(context.Context, game.ID, player.Mark, grid.Position) error
	NextPlayer(context.Context, game.ID) (player.Mark, error)
}
