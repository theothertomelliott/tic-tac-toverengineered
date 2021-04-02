package matchmaker

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
)

// MatchMaker allows a user to request a game, and check on the state
// of an existing request.
type MatchMaker interface {
	// Request submits a new game request asynchronously.
	// The RequestID returned can be submitted to Check.
	Request(context.Context) (RequestID, error)

	// Check returns the current state of a request.
	// If a match has been made, a Match object for a ready game
	// is returned.
	// Otherwise, nil is returned.
	Check(context.Context, RequestID) (*Match, error)
}

type Match struct {
	Game  game.ID
	Mark  string
	Token PlayerToken
}

// RequestID is a unique identifier for a match request.
type RequestID string

// PlayerToken is a unique secret identifying a player who has
// been matched to a game.
// It will be used to authenticate the player when making moves
// in the game.
type PlayerToken string
