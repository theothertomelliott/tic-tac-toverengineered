package inmemoryturns

import (
	"context"
	"fmt"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/checker/pkg/win"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/currentturn/pkg/turn"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid"
	"go.opentelemetry.io/otel"
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

func (c *controller) NextPlayer(ctx context.Context, g game.ID) (player.Mark, error) {
	tracer := otel.GetTracerProvider().Tracer("TurnController")
	ctx, span := tracer.Start(ctx, "NextPlayer")
	defer span.End()

	return c.current.Player(ctx, g)
}

func (c *controller) TakeTurn(ctx context.Context, g game.ID, m player.Mark, p grid.Position) error {
	tracer := otel.GetTracerProvider().Tracer("TurnController")
	ctx, span := tracer.Start(ctx, "TakeTurn")
	defer span.End()

	// Ensure it's this player's turn
	curM, err := c.current.Player(ctx, g)
	if err != nil {
		return fmt.Errorf("could not confirm player turn: %w", err)
	}

	if curM != m {
		return fmt.Errorf("not player %s's turn", m.String())
	}

	// Don't allow play after the game is won
	w, err := c.checker.Winner(ctx, g)
	if err != nil {
		return fmt.Errorf("could not check win status: %w", err)
	}
	if w.Finished() {
		return fmt.Errorf("game is complete")
	}

	// Make the mark
	if err := c.grid.SetMark(ctx, g, p, m); err != nil {
		return fmt.Errorf("could not take turn: %w", err)
	}

	if err := c.current.Next(ctx, g); err != nil {
		return fmt.Errorf("could not advance turn: %w", err)
	}
	return nil
}
