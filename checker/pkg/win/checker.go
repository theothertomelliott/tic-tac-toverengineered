package win

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
)

type Checker interface {
	Winner(context.Context, game.ID) (*player.Mark, error)
}
