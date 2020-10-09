package randombot_test

import (
	"testing"

	"github.com/theothertomelliott/tic-tac-toverengineered/bot/pkg/bot/randombot"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid"
)

func TestBotSelectsAnySpace(t *testing.T) {
	var g = [][]*player.Mark{
		{nil, nil, nil},
		{nil, nil, nil},
		{nil, nil, nil},
	}

	b := randombot.New()
	for i := 0; i < 100; i++ {
		got, err := b.Move(player.X, g)
		if err != nil {
			t.Errorf("test %d> unexpected error: %v", i, err)
		}
		if got.X < 0 || got.X > 2 {
			t.Errorf("test %d> x was out of bounds", i)
		}
		if got.Y < 0 || got.Y > 2 {
			t.Errorf("test %d> y was out of bounds", i)
		}
	}
}

func TestBotSelectsTheOneValidSpace(t *testing.T) {
	x := player.MarkToPointer(player.X)
	var g = [][]*player.Mark{
		{x, x, x},
		{x, x, x},
		{x, nil, x},
	}

	b := randombot.New()
	for i := 0; i < 100; i++ {
		got, err := b.Move(player.X, g)
		if err != nil {
			t.Errorf("test %d> unexpected error: %v", i, err)
		}
		expected := grid.Position{X: 2, Y: 1}
		if got != expected {
			t.Errorf("expected %+v, got %+v", expected, got)
		}
	}
}

func TestBotErrorsWhenNoMove(t *testing.T) {
	x := player.MarkToPointer(player.X)
	var g = [][]*player.Mark{
		{x, x, x},
		{x, x, x},
		{x, x, x},
	}

	b := randombot.New()
	for i := 0; i < 100; i++ {
		_, err := b.Move(player.X, g)
		if err == nil {
			t.Errorf("test %d> expected an error", i)
		}
	}
}
