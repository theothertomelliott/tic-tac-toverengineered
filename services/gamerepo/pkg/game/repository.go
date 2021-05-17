package game

import "context"

// Repository handles creation and tracking of games
type Repository interface {
	// New creates a new game and creates a unique ID
	New(context.Context) (ID, error)
	// List obtains game IDs, ordered by creation date.
	// Pagination is managed through the max and offset params.
	List(ctx context.Context, max int64, offset int64) ([]ID, error)
	// Exists returns true iff the given game ID was previously created with New
	Exists(context.Context, ID) (bool, error)
}
