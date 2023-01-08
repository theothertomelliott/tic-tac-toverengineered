package grid

import (
	"context"
	"fmt"
	"sync"

	"github.com/gammazero/workerpool"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	space "github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg/spaceinmemory"
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
		wp:     workerpool.New(100),
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
	wp     *workerpool.WorkerPool
}

func (g *gridImpl) State(ctx context.Context, gameID game.ID) ([][]*player.Mark, error) {
	var out [][]*player.Mark

	// Fill out the grid
	for _, r := range g.spaces {
		var row []*player.Mark
		for range r {
			row = append(row, nil)
		}
		out = append(out, row)
	}

	var wg sync.WaitGroup
	wg.Add(len(out) * len(out[0]))

	var errors = make(chan error)
	for i, r := range g.spaces {
		for j, s := range r {
			func(i, j int, s space.Space) {
				g.wp.Submit(func() {
					defer wg.Done()
					mark, err := s.Mark(ctx, gameID)
					if err != nil {
						errors <- fmt.Errorf("could not get mark at (%d,%d): %w", i, j, err)
						return
					}
					out[i][j] = mark
				})
			}(i, j, s)
		}
	}
	wg.Wait()

	select {
	case err := <-errors:
		return nil, err
	default:
	}

	return out, nil
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
	if err := g.spaces[p.X][p.Y].SetMark(ctx, game, m); err != nil {
		return fmt.Errorf("%v: %w", p, err)
	}
	return nil
}

func (g *gridImpl) Rows(ctx context.Context) ([]Row, error) {
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
	}, nil
}
