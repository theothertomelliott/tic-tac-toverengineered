package inmemoryrepository

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	"go.opentelemetry.io/otel"
)

// New creates a new game.Repository with
// games recorded in memory.
func New() game.Repository {
	return &repository{
		gameSet: make(map[game.ID]struct{}),
	}
}

type repository struct {
	games   []game.ID
	gameSet map[game.ID]struct{} // lookup table for game IDs
	mtx     sync.RWMutex
}

// New creates a new game and creates a unique ID
func (r *repository) New(ctx context.Context) (game.ID, error) {
	tracer := otel.GetTracerProvider().Tracer("Gamerepo")
	ctx, span := tracer.Start(ctx, "New")
	defer span.End()

	r.mtx.Lock()
	defer r.mtx.Unlock()

	u := uuid.New()
	id := game.ID(u.String())
	r.games = append(r.games, id)
	r.gameSet[id] = struct{}{}
	return id, nil
}

// List obtains game IDs, ordered by creation date.
// Pagination is managed through the max and offset params.
func (r *repository) List(ctx context.Context, max int64, offset int64) ([]game.ID, error) {
	tracer := otel.GetTracerProvider().Tracer("Gamerepo")
	ctx, span := tracer.Start(ctx, "List")
	defer span.End()

	r.mtx.RLock()
	defer r.mtx.RUnlock()

	total := int64(len(r.games))
	if offset >= total {
		return []game.ID{}, nil
	}
	if offset+max >= total {
		return r.games[offset:], nil
	}
	return r.games[offset : offset+max], nil
}

// Exists returns true iff the given game ID was previously created with New
func (r *repository) Exists(ctx context.Context, id game.ID) (bool, error) {
	tracer := otel.GetTracerProvider().Tracer("Gamerepo")
	ctx, span := tracer.Start(ctx, "Exists")
	defer span.End()

	r.mtx.RLock()
	defer r.mtx.RUnlock()

	_, exists := r.gameSet[id]
	return exists, nil
}
