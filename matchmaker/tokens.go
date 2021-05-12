package matchmaker

import (
	"errors"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
)

// InvalidToken indicates that a given token was not a real
// player token.
var InvalidToken = errors.New("invalid player token")

// TokenCreator cuts tokens to be issued to end users.
type TokenCreator interface {
	// Create returns a token for a provided game and player mark.
	Create(game.ID, player.Mark) (string, error)
}

// TokenValidator verifies a provided token
type TokenValidator interface {
	// Validate checks a provided token and returns the
	// associated game id and mark if valid.
	// If not valid, an error of type InvalidToken is returned.
	Validate(token string) (game.ID, player.Mark, error)
}
