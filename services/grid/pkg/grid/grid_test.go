package grid_test

import (
	"context"
	"testing"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid"
	space "github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg/spaceinmemory"
)

const testID = game.ID("test")

// TestNewErrors verifies that grids can only be created
// with a 3x3 array of spaces.
func TestNewErrors(t *testing.T) {
	ns := func() space.Space {
		return spaceinmemory.New()
	}
	var tests = []struct {
		name   string
		spaces [][]space.Space
	}{
		{
			name:   "nil",
			spaces: nil,
		},
		{
			name: "1x1",
			spaces: [][]space.Space{
				{ns()},
			},
		},
		{
			name: "1x3 rectangle",
			spaces: [][]space.Space{
				{ns(), ns(), ns()},
			},
		},
		{
			name: "3x1 rectangle",
			spaces: [][]space.Space{
				{ns()},
				{ns()},
				{ns()},
			},
		},
		{
			name: "irregular array",
			spaces: [][]space.Space{
				{ns()},
				{ns(), ns(), ns()},
				{ns()},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := grid.New(test.spaces)
			if err == nil {
				t.Error("expected an error")
			}
		})
	}
}

func TestMarkSpace(t *testing.T) {
	g := grid.NewInMemory()
	pos := grid.Position{X: 0, Y: 0}
	g.SetMark(context.Background(), testID, pos, player.X)
	gotMark, err := g.Mark(context.Background(), testID, pos)
	if err != nil {
		t.Error(err)
	}
	if gotMark == nil || *gotMark != player.X {
		t.Errorf("expected %v, got %v", player.X, gotMark)
	}
}
func TestCanOnlyMarkSpaceOnce(t *testing.T) {
	g := grid.NewInMemory()
	pos := grid.Position{X: 0, Y: 0}
	if err := g.SetMark(context.Background(), testID, pos, player.X); err != nil {
		t.Error(err)
	}
	if err := g.SetMark(context.Background(), testID, pos, player.X); err == nil {
		t.Error("Expected an error when attempting to mark a space again")
	}
}

func TestRows(t *testing.T) {
	g := grid.NewInMemory()
	sr := g.Rows(context.Background())
	if len(sr) != 8 {
		t.Errorf("expected 8 rows, got %d", len(sr))
	}
	for _, row := range sr {
		if len(row) != 3 {
			t.Errorf("rows should have 3 positions, got %v", row)
		}
	}
}
