package inmemoryrepository

import (
	"github.com/google/uuid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
)

// New creates a new game.Repository with
// games recorded in memory.
func New() game.Repository {
	return &repository{}
}

type repository struct {
	games []game.ID
}

// New creates a new game and creates a unique ID
func (r *repository) New() game.ID {
	u := uuid.New()
	id := game.ID(u.String())
	r.games = append(r.games, id)
	return id
}

// List obtains game IDs, ordered by creation date.
// Pagination is managed through the max and offset params.
func (r *repository) List(max int64, offset int64) []game.ID {
	panic("not implemented") // TODO: Implement
}