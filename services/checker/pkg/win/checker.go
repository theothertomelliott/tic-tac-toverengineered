package win

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
)

type Result struct {
	Winner *player.Mark
	IsDraw bool
}

// Finished returns true iff the game is complete
func (r Result) Finished() bool {
	return r.IsDraw || r.Winner != nil
}

func (r Result) Equal(res Result) bool {
	var aWinner, bWinner string
	if r.Winner != nil {
		aWinner = string(*r.Winner)
	}
	if res.Winner != nil {
		bWinner = string(*res.Winner)
	}
	return aWinner == bWinner && r.IsDraw == res.IsDraw
}

type Checker interface {
	Winner(context.Context, game.ID) (Result, error)
}
