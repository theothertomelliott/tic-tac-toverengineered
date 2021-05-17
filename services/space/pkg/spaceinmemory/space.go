package spaceinmemory

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	space "github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg"
)

var _ space.Space = &Space{}

func New() space.Space {
	return &Space{
		markByGame: make(map[game.ID]*player.Mark),
	}
}

func NewWithMark(g game.ID, m *player.Mark) space.Space {
	return &Space{
		markByGame: map[game.ID]*player.Mark{
			g: m,
		},
	}
}

// Space implements the Space interface in memory
type Space struct {
	markByGame map[game.ID]*player.Mark
}

// Mark returns the mark applied to this space, or nil if there is none
func (s *Space) Mark(ctx context.Context, game game.ID) (*player.Mark, error) {
	return s.markByGame[game], nil
}

// SetMark applies the specified mark to this space.
func (s *Space) SetMark(ctx context.Context, game game.ID, m player.Mark) error {
	s.markByGame[game] = &m
	return nil
}
