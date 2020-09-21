package space

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
)

type Space interface {
	Mark(context.Context, game.ID) (*player.Mark, error)
	SetMark(context.Context, game.ID, player.Mark) error
}
