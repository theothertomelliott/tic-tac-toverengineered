package turn

import "github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"

type Current interface {
	Player() (player.Mark, error)
	Next() error
}
