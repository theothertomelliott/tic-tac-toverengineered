package turn

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
)

type Controller interface {
	TakeTurn(game.ID, player.Mark, grid.Position) error
	NextPlayer(game.ID) (player.Mark, error)
}
