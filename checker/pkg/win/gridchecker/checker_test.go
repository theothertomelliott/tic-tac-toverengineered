package gridchecker_test

import (
	"context"
	"testing"

	"github.com/theothertomelliott/tic-tac-toverengineered/checker/pkg/win/gridchecker"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
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
		name     string
		grid     grid.Grid
		expected *player.Mark
	}{
		{
			name:     "empty grid",
			grid:     grid.NewInMemory(),
			expected: nil,
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
			expected: player.MarkToPointer(player.X),
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
			expected: player.MarkToPointer(player.O),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gridchecker.New(test.grid)
			got, err := c.Winner(context.Background(), testID)
			if err != nil {
				t.Error(err)
			}
			if got.String() != test.expected.String() {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
