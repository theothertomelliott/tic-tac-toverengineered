package win

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
)

type Checker interface {
	Winner(game.ID) (*player.Mark, error)
}
