package turn

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
)

type Current interface {
	Player(game.ID) (player.Mark, error)
	Next(game.ID) error
}
