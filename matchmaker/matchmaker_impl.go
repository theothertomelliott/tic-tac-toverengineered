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
) MatchMakerServer {
	return &matchMaker{
		matches: matches,
		queue:   queue,
		games:   games,
	}
}

var _ MatchMakerServer = &matchMaker{}

// RequestQueue defines a queue of match requests.
//
// This queue only needs to be able to hold one request at a time, as
// any subsequent request will result in a match between the new request
//  and the request already in the queue.
type RequestQueue interface {
	// Enqueue adds a request to the queue
	Enqueue(context.Context, string) error
	// Dequeue removes a request from the queue
	Dequeue(context.Context) (*string, error)
}

// MatchStore holds match results associated with the originating
// requests.
type MatchStore interface {
	// Set associates a request id with a match result
	Set(context.Context, string, Match) error
	// Get retrieves a match result for a request id.
	// If no match exists, nil is returned.
	Get(context.Context, string) (*Match, error)
}

type matchMaker struct {
	matches MatchStore
	queue   RequestQueue
	games   game.Repository
}

func (m *matchMaker) Request(ctx context.Context, req *RequestRequest) (*RequestResponse, error) {
	id := uuid.New().String()
	partner, err := m.queue.Dequeue(ctx)
	if err != nil {
		return nil, err
	}
	if partner != nil {
		game, err := m.games.New(ctx)
		if err != nil {
			return nil, err
		}

		m.matches.Set(ctx, id, Match{
			GameId: string(game),
			Mark:   "O",
			Token:  uuid.New().String(),
		})
		m.matches.Set(ctx, *partner, Match{
			GameId: string(game),
			Mark:   "X",
			Token:  uuid.New().String(),
		})
	} else {
		err := m.queue.Enqueue(ctx, id)
		if err != nil {
			return nil, err
		}
	}
	return &RequestResponse{
		RequestId: id,
	}, nil
}

func (m *matchMaker) Check(ctx context.Context, req *CheckRequest) (*CheckResponse, error) {
	match, err := m.matches.Get(ctx, req.RequestId)
	return &CheckResponse{
		Match: match,
	}, err
}
