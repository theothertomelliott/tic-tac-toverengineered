package game

// Repository handles creation and tracking of games
type Repository interface {
	// New creates a new game and creates a unique ID
	New() (ID, error)
	// List obtains game IDs, ordered by creation date.
	// Pagination is managed through the max and offset params.
	List(max int64, offset int64) ([]ID, error)
}
