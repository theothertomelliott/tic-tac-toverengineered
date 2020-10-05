package turn

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
)

type Current interface {
	Player(context.Context, game.ID) (player.Mark, error)
	Next(context.Context, game.ID) error
}
