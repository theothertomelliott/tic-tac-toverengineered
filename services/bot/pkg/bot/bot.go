package bot

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
)

// Bot defines an interface for a computer Tic-Tac-Toe player.
type Bot interface {
	Move(mark player.Mark, state [][]string) (tictactoeapi.Position, error)
}
