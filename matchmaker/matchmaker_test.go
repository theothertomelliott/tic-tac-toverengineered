package matchmaker_test

import (
	"context"
	"sync"
	"testing"

	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game/inmemoryrepository"
	"github.com/theothertomelliott/tic-tac-toverengineered/matchmaker"
)

func TestSingleMatch(t *testing.T) {
	games := inmemoryrepository.New()
	m := matchmaker.New(games, newQueue(), newStore())

	doTestMatch(t, m)
}

func TestMultipleMatches(t *testing.T) {
	games := inmemoryrepository.New()
	m := matchmaker.New(games, newQueue(), newStore())

	for i := 0; i < 100; i++ {
		doTestMatch(t, m)
	}
}

func doTestMatch(t *testing.T, m matchmaker.MatchMaker) {
	// First player
	p1Req, err := m.Request(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	match1, err := m.Check(context.Background(), p1Req)
	if err != nil {
		t.Fatal(err)
	}
	if match1 != nil {
		t.Errorf("expected no match, got: %v", match1)
	}

	// Second player
	p2Req, err := m.Request(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if p2Req == p1Req {
		t.Errorf("request IDs must be different")
	}
	match2, err := m.Check(context.Background(), p2Req)
	if err != nil {
		t.Fatal(err)
	}
	if match2 == nil {
		t.Errorf("expected a match on second request")
	}
	if match2.Mark != "O" {
		t.Errorf("second player should be O, got: %v", match2.Mark)
	}

	// Poll first player match
	match3, err := m.Check(context.Background(), p1Req)
	if err != nil {
		t.Fatal(err)
	}
	if match3 == nil {
		t.Errorf("expected a match on third request")
	}
	if match3.Mark != "X" {
		t.Errorf("first player should be X, got: %v", match3.Mark)
	}

	if match2.Game != match3.Game {
		t.Errorf("game ids must match")
	}

	if match2.Token == match3.Token {
		t.Errorf("player tokens must not match")
	}
}

var _ matchmaker.RequestQueue = &channelRequestQueue{}

type channelRequestQueue struct {
	requests chan matchmaker.RequestID
}

func newQueue() matchmaker.RequestQueue {
	return &channelRequestQueue{
		requests: make(chan matchmaker.RequestID, 1),
	}
}

var _ matchmaker.MatchStore = &matchStore{}

type matchStore struct {
	mtx     sync.Mutex
	matches map[matchmaker.RequestID]*matchmaker.Match
}

func newStore() matchmaker.MatchStore {
	return &matchStore{
		matches: make(map[matchmaker.RequestID]*matchmaker.Match),
	}
}

func (m *matchStore) Set(ctx context.Context, req matchmaker.RequestID, match matchmaker.Match) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.matches[req] = &match
	return nil
}

func (m *matchStore) Get(ctx context.Context, req matchmaker.RequestID) (*matchmaker.Match, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	return m.matches[req], nil
}

func (c *channelRequestQueue) Enqueue(_ context.Context, id matchmaker.RequestID) error {
	c.requests <- id
	return nil
}

func (c *channelRequestQueue) Dequeue(_ context.Context) (*matchmaker.RequestID, error) {
	select {
	case id := <-c.requests:
		return &id, nil
	default:
		return nil, nil
	}
}
