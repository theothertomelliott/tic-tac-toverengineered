package inmemorymatchmaker

import (
	"context"
	"sync"

	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/matchmaker"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/matchmaker/server"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/matchmaker/unsignedtokens"
	"google.golang.org/grpc"
)

var _ matchmaker.MatchMakerClient = &client{}

func New(games game.Repository) matchmaker.MatchMakerClient {
	return &client{
		server: server.New(games, newQueue(), newStore(), &unsignedtokens.UnsignedTokens{}),
	}
}

type client struct {
	server matchmaker.MatchMakerServer
}

func (c *client) Request(ctx context.Context, in *matchmaker.RequestRequest, opts ...grpc.CallOption) (*matchmaker.RequestResponse, error) {
	return c.server.Request(ctx, in)
}

func (c *client) Check(ctx context.Context, in *matchmaker.CheckRequest, opts ...grpc.CallOption) (*matchmaker.CheckResponse, error) {
	return c.server.Check(ctx, in)
}

var _ server.RequestQueue = &channelRequestQueue{}

type channelRequestQueue struct {
	requests chan string
}

func newQueue() server.RequestQueue {
	return &channelRequestQueue{
		requests: make(chan string, 1),
	}
}

var _ server.MatchStore = &matchStore{}

type matchStore struct {
	mtx     sync.Mutex
	matches map[string]*matchmaker.Match
}

func newStore() server.MatchStore {
	return &matchStore{
		matches: make(map[string]*matchmaker.Match),
	}
}

func (m *matchStore) Set(ctx context.Context, req string, match *matchmaker.Match) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.matches[req] = match
	return nil
}

func (m *matchStore) Get(ctx context.Context, req string) (*matchmaker.Match, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	return m.matches[req], nil
}

func (c *channelRequestQueue) Enqueue(_ context.Context, id string) error {
	c.requests <- id
	return nil
}

func (c *channelRequestQueue) Dequeue(_ context.Context) (*string, error) {
	select {
	case id := <-c.requests:
		return &id, nil
	default:
		return nil, nil
	}
}
