package bot

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
)

type Position struct {
	I int32
	J int32
}

// Bot defines an interface for a computer Tic-Tac-Toe player.
type Bot interface {
	Move(mark player.Mark, state [][]string) (Position, error)
}
