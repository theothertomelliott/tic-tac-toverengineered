package grid

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
)

// Grid provides functions to interact with a grid
type Grid interface {
	Mark(Position) (*player.Mark, error)
	SetMark(Position, player.Mark) error
	Rows() []Row
}

// Position defines a position in the grid
type Position struct {
	X int
	Y int
}

// Row represents a row of spaces by their position on the grid
type Row []Position
