package spaceinmemory

import (
	"context"
	"sync"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	space "github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg"
	"go.opentelemetry.io/otel"
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
	mtx        sync.RWMutex
}

// Mark returns the mark applied to this space, or nil if there is none
func (s *Space) Mark(ctx context.Context, game game.ID) (*player.Mark, error) {
	tracer := otel.GetTracerProvider().Tracer("Space")
	ctx, span := tracer.Start(ctx, "Mark")
	defer span.End()

	s.mtx.RLock()
	defer s.mtx.RUnlock()

	return s.markByGame[game], nil
}

// SetMark applies the specified mark to this space.
func (s *Space) SetMark(ctx context.Context, game game.ID, m player.Mark) error {
	tracer := otel.GetTracerProvider().Tracer("Space")
	ctx, span := tracer.Start(ctx, "SetMark")
	defer span.End()

	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.markByGame[game] = &m
	return nil
}
