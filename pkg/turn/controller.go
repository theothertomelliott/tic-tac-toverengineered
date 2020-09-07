package turn

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
)

type Controller interface {
	TakeTurn(player.Mark, grid.Position) error
	NextPlayer() (player.Mark, error)
}
