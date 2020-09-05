package grid

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/space"
)

// Grid provides functions to interact with a grid
type Grid interface {
	Mark(Position) (*player.Mark, error)
	SetMark(Position, player.Mark) error
	PositionRows() []PositionRow
	SpaceRows() []SpaceRow
}

// Position defines a position in the grid
type Position struct {
	X int
	Y int
}

// PositionRow represents a row of spaces by their position on the grid
type PositionRow []Position

// SpaceRow represents a row of spaces by the instances of space.Space representing them
type SpaceRow []space.Space
