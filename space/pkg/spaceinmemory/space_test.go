package spaceinmemory_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/space/pkg/spaceinmemory"
)

func TestSpaceCanBeWritten(t *testing.T) {
	testID := game.ID("test")
	s := spaceinmemory.New()
	got, err := s.Mark(context.Background(), testID)
	if err != nil {
		t.Error(err)
	}
	if got != nil {
		t.Errorf("expected empty space on initialization, got %s", got.String())
	}

	if err := s.SetMark(context.Background(), testID, player.X); err != nil {
		t.Error(err)
	}

	got, err = s.Mark(context.Background(), testID)
	if err != nil {
		t.Error(err)
	}
	if got == nil || *got != player.X {
		t.Errorf("expected empty space on initialization, got %s", got.String())
	}
}

func TestSpacesHandleSeparateGames(t *testing.T) {
	firstID := game.ID("game1")
	secondID := game.ID("game2")
	s := spaceinmemory.New()
	if err := s.SetMark(context.Background(), firstID, player.X); err != nil {
		t.Error(err)
	}
	if err := s.SetMark(context.Background(), secondID, player.O); err != nil {
		t.Error(err)
	}

	gotFirst, err := s.Mark(context.Background(), firstID)
	if err != nil {
		t.Error(err)
	}
	gotSecond, err := s.Mark(context.Background(), secondID)
	if err != nil {
		t.Error(err)
	}

	if fmt.Sprint(gotFirst) != "X" {
		t.Errorf("game 1: expected X, got %v", gotFirst)
	}
	if fmt.Sprint(gotSecond) != "O" {
		t.Errorf("game 2: expected O, got %v", gotFirst)
	}
}
