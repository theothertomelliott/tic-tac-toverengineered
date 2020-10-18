package rpcchecker

import (
	context "context"

	"github.com/theothertomelliott/tic-tac-toverengineered/checker/pkg/win"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
	grpc "google.golang.org/grpc"
)

func ConnectChecker(target string) (*Checker, error) {
	var err error
	c := &Checker{}
	c.conn, err = grpc.Dial(
		target,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(monitoring.UnaryClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}
	c.client = NewCheckerClient(c.conn)
	return c, nil
}

type Checker struct {
	conn   *grpc.ClientConn
	client CheckerClient
}

func (c *Checker) Winner(ctx context.Context, id game.ID) (win.Result, error) {
	resp, err := c.client.Winner(ctx, &WinnerRequest{
		GameId: string(id),
	})
	if err != nil {
		return win.Result{}, err
	}
	if !resp.HasWinner {
		return win.Result{
			IsDraw: resp.IsDraw,
		}, nil
	}
	m := player.Mark(resp.Mark)

	return win.Result{
		Winner: &m,
	}, nil
}
