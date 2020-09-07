package space

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
)

type Space interface {
	Mark(game game.ID) (*player.Mark, error)
	SetMark(game.ID, player.Mark) error
}
