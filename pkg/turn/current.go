package turn

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
)

type Current interface {
	Player(context.Context, game.ID) (player.Mark, error)
	Next(context.Context, game.ID) error
}
