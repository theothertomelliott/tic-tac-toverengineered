package inmemoryturns

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/currentturn/pkg/turn"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
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

func (c *current) Player(ctx context.Context, g game.ID) (player.Mark, error) {
	m, exists := c.mark[g]
	if !exists {
		m = player.X
		c.mark[g] = m
	}
	return m, nil
}

func (c *current) Next(ctx context.Context, g game.ID) error {
	prev, err := c.Player(ctx, g)
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
