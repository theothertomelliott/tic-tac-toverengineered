package bot

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid"
)

// Bot defines an interface for a computer Tic-Tac-Toe player.
type Bot interface {
	Move(mark player.Mark, state [][]*player.Mark) (grid.Position, error)
}
