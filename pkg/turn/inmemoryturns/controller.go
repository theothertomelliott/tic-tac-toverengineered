package inmemoryturns

import (
	"fmt"

	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/turn"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/win"
)

func New(
	current turn.Current,
	grid grid.Grid,
	checker win.Checker,
) turn.Controller {
	return &controller{
		current: current,
		grid:    grid,
		checker: checker,
	}
}

type controller struct {
	current turn.Current
	grid    grid.Grid
	checker win.Checker
}

func (c *controller) NextPlayer(g game.ID) (player.Mark, error) {
	return c.current.Player(g)
}

func (c *controller) TakeTurn(g game.ID, m player.Mark, p grid.Position) error {
	// Ensure it's this player's turn
	curM, err := c.current.Player(g)
	if err != nil {
		return fmt.Errorf("could not confirm player turn: %w", err)
	}

	if curM != m {
		return fmt.Errorf("not player %s's turn", m.String())
	}

	// Don't allow play after the game is won
	w, err := c.checker.Winner(g)
	if err != nil {
		return fmt.Errorf("could not check win status: %w", err)
	}
	if w != nil {
		return fmt.Errorf("game was already won by %v", w)
	}

	// Make the mark
	if err := c.grid.SetMark(g, p, m); err != nil {
		return fmt.Errorf("could not take turn: %w", err)
	}

	if err := c.current.Next(g); err != nil {
		return fmt.Errorf("could not advance turn: %w", err)
	}
	return nil
}
