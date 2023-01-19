package server

import (
	"context"

	"github.com/google/uuid"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/matchmaker"
	"go.opentelemetry.io/otel"
)

func New(
	games game.Repository,
	queue RequestQueue,
	matches MatchStore,
	tokens matchmaker.TokenCreator,
) matchmaker.MatchMakerServer {
	return &matchMaker{
		matches: matches,
		queue:   queue,
		games:   games,
		tokens:  tokens,
	}
}

var _ matchmaker.MatchMakerServer = &matchMaker{}

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
	Set(context.Context, string, *matchmaker.Match) error
	// Get retrieves a match result for a request id.
	// If no match exists, nil is returned.
	Get(context.Context, string) (*matchmaker.Match, error)
}

type matchMaker struct {
	matchmaker.UnimplementedMatchMakerServer

	matches MatchStore
	queue   RequestQueue
	games   game.Repository
	tokens  matchmaker.TokenCreator
}

func (m *matchMaker) RequestPair(ctx context.Context, req *matchmaker.RequestPairRequest) (*matchmaker.RequestPairResponse, error) {
	tracer := otel.GetTracerProvider().Tracer("MatchMaker")
	ctx, span := tracer.Start(ctx, "RequestPair")
	defer span.End()

	game, err := m.games.New(ctx)
	if err != nil {
		return nil, err
	}

	var out = &matchmaker.RequestPairResponse{}

	tokenO, err := m.tokens.Create(game, player.O)
	if err != nil {
		return nil, err
	}
	out.O = &matchmaker.Match{
		GameId: string(game),
		Mark:   "O",
		Token:  tokenO,
	}

	tokenX, err := m.tokens.Create(game, player.X)
	if err != nil {
		return nil, err
	}
	out.X = &matchmaker.Match{
		GameId: string(game),
		Mark:   "X",
		Token:  tokenX,
	}

	return out, nil
}

func (m *matchMaker) Request(ctx context.Context, req *matchmaker.RequestRequest) (*matchmaker.RequestResponse, error) {
	tracer := otel.GetTracerProvider().Tracer("MatchMaker")
	ctx, span := tracer.Start(ctx, "Request")
	defer span.End()

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

		tokenO, err := m.tokens.Create(game, player.O)
		if err != nil {
			return nil, err
		}
		m.matches.Set(ctx, id, &matchmaker.Match{
			GameId: string(game),
			Mark:   "O",
			Token:  tokenO,
		})

		tokenX, err := m.tokens.Create(game, player.X)
		if err != nil {
			return nil, err
		}
		m.matches.Set(ctx, *partner, &matchmaker.Match{
			GameId: string(game),
			Mark:   "X",
			Token:  tokenX,
		})
	} else {
		err := m.queue.Enqueue(ctx, id)
		if err != nil {
			return nil, err
		}
	}
	return &matchmaker.RequestResponse{
		RequestId: id,
	}, nil
}

func (m *matchMaker) Check(ctx context.Context, req *matchmaker.CheckRequest) (*matchmaker.CheckResponse, error) {
	tracer := otel.GetTracerProvider().Tracer("MatchMaker")
	ctx, span := tracer.Start(ctx, "Check")
	defer span.End()

	match, err := m.matches.Get(ctx, req.RequestId)
	return &matchmaker.CheckResponse{
		Match: match,
	}, err
}
