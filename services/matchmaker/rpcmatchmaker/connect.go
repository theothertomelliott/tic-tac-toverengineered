package rpcmatchmaker

import (
	"github.com/theothertomelliott/tic-tac-toverengineered/services/matchmaker"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

// Connect establishes a connection to a repository server and returns a
// client.
func Connect(target string) (matchmaker.MatchMakerClient, error) {
	conn, err := grpc.Dial(
		target,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}
	return matchmaker.NewMatchMakerClient(conn), nil
}
