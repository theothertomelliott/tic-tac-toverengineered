package space

import "github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"

type Space interface {
	Mark() (*player.Mark, error)
	SetMark(player.Mark) error
}
