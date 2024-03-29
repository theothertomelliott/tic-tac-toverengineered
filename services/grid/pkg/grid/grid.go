package grid

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
)

// Grid provides functions to interact with a grid
type Grid interface {
	Mark(context.Context, game.ID, Position) (*player.Mark, error)
	State(context.Context, game.ID) ([][]*player.Mark, error)
	SetMark(context.Context, game.ID, Position, player.Mark) error
	Rows(context.Context) ([]Row, error)
}

// Position defines a position in the grid
type Position struct {
	X int
	Y int
}

// Row represents a row of spaces by their position on the grid
type Row []Position
