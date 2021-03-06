package inmemoryturns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/checker/pkg/win/gridchecker"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/currentturn/pkg/turn/inmemoryturns"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid"
	space "github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg/spaceinmemory"
)

const testID = game.ID("test")

func TestControllerAppliesMark(t *testing.T) {
	g := grid.NewInMemory()
	c := inmemoryturns.New(
		inmemoryturns.NewCurrentTurn(),
		g,
		gridchecker.New(g),
	)
	pos := grid.Position{X: 0, Y: 0}

	// X goes first
	if err := c.TakeTurn(context.Background(), testID, player.X, pos); err != nil {
		t.Fatal(err)
	}
	m, err := g.Mark(context.Background(), testID, pos)
	if err != nil {
		t.Error(err)
	}
	got := fmt.Sprint(m)
	expected := fmt.Sprint(player.MarkToPointer(player.X))
	if expected != got {
		t.Errorf("expected %q, got %q", expected, got)
	}

}

func TestControllerTurns(t *testing.T) {
	g := grid.NewInMemory()
	c := inmemoryturns.New(
		inmemoryturns.NewCurrentTurn(),
		g,
		gridchecker.New(g),
	)
	// X goes first
	if err := c.TakeTurn(context.Background(), testID, player.X, grid.Position{X: 0, Y: 0}); err != nil {
		t.Fatal(err)
	}

	// X cannot go again
	if err := c.TakeTurn(context.Background(), testID, player.X, grid.Position{X: 0, Y: 0}); err == nil {
		t.Fatal("expected an error")
	}

	// O goes second
	if err := c.TakeTurn(context.Background(), testID, player.O, grid.Position{X: 1, Y: 0}); err != nil {
		t.Fatal(err)
	}
}

func TestCannotPlayAfterWin(t *testing.T) {
	x := func() space.Space {
		m := player.X
		return spaceinmemory.NewWithMark(testID, &m)
	}
	o := func() space.Space {
		m := player.O
		return spaceinmemory.NewWithMark(testID, &m)
	}
	g, _ := grid.New([][]space.Space{
		{x(), x(), o()},
		{x(), o(), o()},
		{o(), o(), x()},
	})

	c := inmemoryturns.New(
		inmemoryturns.NewCurrentTurn(),
		g,
		gridchecker.New(g),
	)
	// X cannot play
	if err := c.TakeTurn(context.Background(), testID, player.X, grid.Position{X: 0, Y: 0}); err == nil {
		t.Fatal("expected an error")
	}
	// O cannot play
	if err := c.TakeTurn(context.Background(), testID, player.O, grid.Position{X: 0, Y: 0}); err == nil {
		t.Fatal("expected an error")
	}
}
