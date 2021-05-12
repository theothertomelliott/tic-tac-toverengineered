package main

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/defaultmonitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/rpc/rpcui/rpcserver"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game/rpcrepository/repoclient"
	"github.com/theothertomelliott/tic-tac-toverengineered/matchmaker"
	"github.com/theothertomelliott/tic-tac-toverengineered/matchmaker/server"
	"github.com/theothertomelliott/tic-tac-toverengineered/matchmaker/unsignedtokens"
)

func main() {
	version.Println()
	defaultmonitoring.Init("matchmaker")
	defer monitoring.Close()

	port := 8080
	grpcuiPort := 8081

	games, err := repoclient.Connect(getRepoServerTarget())
	if err != nil {
		log.Fatalf("could not connect to repo server: %v", err)
	}

	rpcServer := rpcserver.New(port)
	matchmaker.RegisterMatchMakerServer(rpcServer.GRPC(),
		server.New(
			games,
			newQueue(),
			newStore(),
			&unsignedtokens.UnsignedTokens{},
		),
	)

	log.Printf("gRPC listening on port :%v", port)
	var done = make(chan struct{})
	go func() {
		err := rpcServer.Serve()
		if err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		time.Sleep(time.Second)
		err := rpcui.Start(port, grpcuiPort)
		if err != nil {
			log.Printf("Failed to start gRPCUI: %v", err)
		}
	}()
	<-done

}

func getRepoServerTarget() string {
	if serverTarget := os.Getenv("REPO_SERVER_TARGET"); serverTarget != "" {
		return serverTarget
	}
	return "localhost:8082"
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

func (m *matchStore) Set(ctx context.Context, req string, match matchmaker.Match) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.matches[req] = &match
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
