package gridchecker

import (
	"fmt"

	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/win"
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

func (c *checker) Winner() (*player.Mark, error) {
	rows := c.grid.Rows()
	for _, row := range rows {
		var markCounts = make(map[player.Mark]int)
		for _, pos := range row {
			m, err := c.grid.Mark(pos)
			if err != nil {
				return nil, fmt.Errorf("%v: %w", pos, err)
			}
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
