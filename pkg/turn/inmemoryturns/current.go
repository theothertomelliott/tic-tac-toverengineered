package inmemoryturns

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/turn"
)

// NewCurrentTurn creates an in-memory instance of turn.Current
func NewCurrentTurn() turn.Current {
	return &current{
		mark: player.X,
	}
}

type current struct {
	mark player.Mark
}

func (c *current) Player() (player.Mark, error) {
	return c.mark, nil
}

func (c *current) Next() error {
	if c.mark == player.X {
		c.mark = player.O
		return nil
	}
	c.mark = player.X
	return nil
}
