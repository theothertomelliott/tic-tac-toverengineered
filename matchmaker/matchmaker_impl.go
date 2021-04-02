package matchmaker

import (
	"context"

	"github.com/google/uuid"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
)

func New(
	games game.Repository,
	queue RequestQueue,
	matches MatchStore,
) MatchMaker {
	return &matchMaker{
		matches: matches,
		queue:   queue,
		games:   games,
	}
}

var _ MatchMaker = &matchMaker{}

// RequestQueue defines a queue of match requests.
//
// This queue only needs to be able to hold one request at a time, as
// any subsequent request will result in a match between the new request
//  and the request already in the queue.
type RequestQueue interface {
	// Enqueue adds a request to the queue
	Enqueue(context.Context, RequestID) error
	// Dequeue removes a request from the queue
	Dequeue(context.Context) (*RequestID, error)
}

// MatchStore holds match results associated with the originating
// requests.
type MatchStore interface {
	// Set associates a request id with a match result
	Set(context.Context, RequestID, Match) error
	// Get retrieves a match result for a request id.
	// If no match exists, nil is returned.
	Get(context.Context, RequestID) (*Match, error)
}

type matchMaker struct {
	matches MatchStore
	queue   RequestQueue
	games   game.Repository
}

func (m *matchMaker) Request(ctx context.Context) (RequestID, error) {
	id := RequestID(uuid.New().String())
	partner, err := m.queue.Dequeue(ctx)
	if err != nil {
		return "", err
	}
	if partner != nil {
		game, err := m.games.New(ctx)
		if err != nil {
			return "", err
		}

		m.matches.Set(ctx, id, Match{
			Game:  game,
			Mark:  "O",
			Token: PlayerToken(uuid.New().String()),
		})
		m.matches.Set(ctx, *partner, Match{
			Game:  game,
			Mark:  "X",
			Token: PlayerToken(uuid.New().String()),
		})
	} else {
		err := m.queue.Enqueue(ctx, id)
		if err != nil {
			return "", err
		}
	}
	return id, nil
}

func (m *matchMaker) Check(ctx context.Context, request RequestID) (*Match, error) {
	return m.matches.Get(ctx, request)
}
