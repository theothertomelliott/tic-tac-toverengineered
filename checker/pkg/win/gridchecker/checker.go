package gridchecker

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/checker/pkg/win"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid"
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

func (c *checker) Winner(ctx context.Context, game game.ID) (*player.Mark, error) {
	rows := c.grid.Rows(ctx)
	state, err := c.grid.State(ctx, game)
	if err != nil {
		return nil, err
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
				return &mark, nil
			}
		}
	}
	return nil, nil
}
