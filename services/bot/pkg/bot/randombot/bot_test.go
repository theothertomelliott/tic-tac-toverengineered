package randombot_test

import (
	"testing"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/bot/pkg/bot/randombot"
)

func TestBotSelectsAnySpace(t *testing.T) {
	var g = [][]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}

	b := randombot.New()
	for i := 0; i < 100; i++ {
		got, err := b.Move(player.X, g)
		if err != nil {
			t.Errorf("test %d> unexpected error: %v", i, err)
		}
		if got.I < 0 || got.I > 2 {
			t.Errorf("test %d> x was out of bounds", i)
		}
		if got.J < 0 || got.J > 2 {
			t.Errorf("test %d> y was out of bounds", i)
		}
	}
}

func TestBotSelectsTheOneValidSpace(t *testing.T) {
	x := string(player.X)
	var g = [][]string{
		{x, x, x},
		{x, x, x},
		{x, "", x},
	}

	b := randombot.New()
	for i := 0; i < 100; i++ {
		got, err := b.Move(player.X, g)
		if err != nil {
			t.Errorf("test %d> unexpected error: %v", i, err)
		}
		expected := tictactoeapi.Position{I: 2, J: 1}
		if got != expected {
			t.Errorf("expected %+v, got %+v", expected, got)
		}
	}
}

func TestBotErrorsWhenNoMove(t *testing.T) {
	x := string(player.X)
	var g = [][]string{
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
