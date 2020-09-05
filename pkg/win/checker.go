package win

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
)

type Checker interface {
	Winner() *player.Mark
}
