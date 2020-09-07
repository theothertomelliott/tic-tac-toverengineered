package inmemoryturns

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/turn"
)

// NewCurrentTurn creates an in-memory instance of turn.Current
func NewCurrentTurn() turn.Current {
	return &current{
		mark: make(map[game.ID]player.Mark),
	}
}

type current struct {
	mark map[game.ID]player.Mark
}

func (c *current) Player(g game.ID) (player.Mark, error) {
	m, exists := c.mark[g]
	if !exists {
		m = player.X
		c.mark[g] = m
	}
	return m, nil
}

func (c *current) Next(g game.ID) error {
	prev, err := c.Player(g)
	if err != nil {
		return err
	}
	if prev == player.X {
		c.mark[g] = player.O
		return nil
	}
	c.mark[g] = player.X
	return nil
}
