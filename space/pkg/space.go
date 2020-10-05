package space

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
)

type Space interface {
	Mark(context.Context, game.ID) (*player.Mark, error)
	SetMark(context.Context, game.ID, player.Mark) error
}
