package gridchecker_test

import (
	"context"
	"testing"

	"github.com/theothertomelliott/tic-tac-toverengineered/checker/pkg/win"
	"github.com/theothertomelliott/tic-tac-toverengineered/checker/pkg/win/gridchecker"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid"
	space "github.com/theothertomelliott/tic-tac-toverengineered/space/pkg"
	"github.com/theothertomelliott/tic-tac-toverengineered/space/pkg/spaceinmemory"
)

const testID = game.ID("test")

func TestChecker(t *testing.T) {
	x := func() space.Space {
		m := player.X
		return spaceinmemory.NewWithMark(testID, &m)
	}
	o := func() space.Space {
		m := player.O
		return spaceinmemory.NewWithMark(testID, &m)
	}
	var tests = []struct {
		name           string
		grid           grid.Grid
		expected       win.Result
		expectFinished bool
	}{
		{
			name:     "empty grid",
			grid:     grid.NewInMemory(),
			expected: win.Result{},
		},
		{
			name: "x wins",
			grid: func() grid.Grid {
				g, _ := grid.New([][]space.Space{
					{x(), x(), x()},
					{x(), o(), o()},
					{o(), o(), x()},
				})
				return g
			}(),
			expected: win.Result{
				Winner: player.MarkToPointer(player.X),
			},
		},
		{
			name: "diagonal o",
			grid: func() grid.Grid {
				g, _ := grid.New([][]space.Space{
					{x(), x(), o()},
					{x(), o(), o()},
					{o(), o(), x()},
				})
				return g
			}(),
			expected: win.Result{
				Winner: player.MarkToPointer(player.O),
			},
		},
		{
			name: "draw",
			grid: func() grid.Grid {
				g, _ := grid.New([][]space.Space{
					{o(), x(), o()},
					{x(), x(), o()},
					{o(), o(), x()},
				})
				return g
			}(),
			expected: win.Result{
				IsDraw: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gridchecker.New(test.grid)
			got, err := c.Winner(context.Background(), testID)
			if err != nil {
				t.Error(err)
			}
			if !test.expected.Equal(got) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
