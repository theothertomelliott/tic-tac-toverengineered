package gridchecker

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/checker/pkg/win"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid"
)

// New creates a Checker for a given grid
func New(g grid.Grid) win.Checker {
	return &checker{
		grid: g,
	}
}

type checker struct {
	grid grid.Grid
}

func (c *checker) Winner(ctx context.Context, game game.ID) (win.Result, error) {
	rows, err := c.grid.Rows(ctx)
	if err != nil {
		return win.Result{}, err
	}

	state, err := c.grid.State(ctx, game)
	if err != nil {
		return win.Result{}, err
	}
	for _, row := range rows {
		var markCounts = make(map[player.Mark]int)
		for _, pos := range row {
			m := state[pos.X][pos.Y]
			if m == nil {
				continue
			}
			markCounts[*m]++
		}
		for mark, count := range markCounts {
			if count == 3 {
				return win.Result{
					Winner: &mark,
				}, nil
			}
		}
	}

	// If no winner, check for any blank spaces.
	// Any blank spaces indicate that the game is not complete.
	for _, row := range state {
		for _, cell := range row {
			if cell == nil {
				return win.Result{}, nil
			}
		}
	}
	// No blank spaces indicates a draw
	return win.Result{
		IsDraw: true,
	}, nil
}
