package win

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
)

type Checker interface {
	Winner(grid.Grid) *player.Mark
}

type RowChecker interface {
	Winner(grid.SpaceRow) *player.Mark
}
