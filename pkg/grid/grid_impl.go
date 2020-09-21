package grid

import (
	"context"
	"fmt"

	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/space"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/space/spaceinmemory"
)

// New creates a Grid from a given array of spaces.
// The provided array must be 3x3.
func New(spaces [][]space.Space) (Grid, error) {
	// Verify this is a 3x3 square grid
	if len(spaces) != 3 {
		return nil, fmt.Errorf("expected a 3x3 array of spaces")
	}
	for _, r := range spaces {
		if len(r) != 3 {
			return nil, fmt.Errorf("expected a 3x3 array of spaces")
		}
	}
	return &gridImpl{
		spaces: spaces,
	}, nil
}

// NewInMemory creates an empty 3x3 Grid with spaces stored in memory.
func NewInMemory() Grid {
	ns := func() space.Space {
		return spaceinmemory.New()
	}
	g, _ := New([][]space.Space{
		{ns(), ns(), ns()},
		{ns(), ns(), ns()},
		{ns(), ns(), ns()},
	})
	return g
}

type gridImpl struct {
	spaces [][]space.Space
}

func (g *gridImpl) Mark(ctx context.Context, game game.ID, p Position) (*player.Mark, error) {
	m, err := g.spaces[p.X][p.Y].Mark(ctx, game)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", p, err)
	}
	return m, nil
}

func (g *gridImpl) SetMark(ctx context.Context, game game.ID, p Position, m player.Mark) error {
	if existing, err := g.Mark(ctx, game, p); err != nil {
		return fmt.Errorf("could not confirm space was not marked: %w", err)
	} else if existing != nil {
		return fmt.Errorf("%v: space has already been marked", p)
	}
	if err := g.spaces[p.X][p.Y].SetMark(context.TODO(), game, m); err != nil {
		return fmt.Errorf("%v: %w", p, err)
	}
	return nil
}

func (g *gridImpl) Rows(ctx context.Context) []Row {
	p := func(x, y int) Position {
		return Position{
			X: x,
			Y: y,
		}
	}
	return []Row{
		// Horizontal
		{p(0, 0), p(1, 0), p(2, 0)},
		{p(0, 1), p(1, 1), p(2, 1)},
		{p(0, 2), p(1, 2), p(2, 2)},

		// Vertical
		{p(0, 0), p(0, 1), p(0, 2)},
		{p(1, 0), p(1, 1), p(1, 2)},
		{p(2, 0), p(2, 1), p(2, 2)},

		// Diagonal
		{p(0, 0), p(1, 1), p(2, 2)},
		{p(0, 2), p(1, 1), p(2, 0)},
	}
}
