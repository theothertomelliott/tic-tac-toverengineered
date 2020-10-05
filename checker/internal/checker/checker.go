package checker

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/checker/pkg/win"
	"github.com/theothertomelliott/tic-tac-toverengineered/checker/pkg/win/rpcchecker"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
)

func NewServer(checker win.Checker) rpcchecker.CheckerServer {
	return &checkerServer{
		checker: checker,
	}
}

type checkerServer struct {
	checker win.Checker
}

func (c *checkerServer) Winner(ctx context.Context, req *rpcchecker.WinnerRequest) (*rpcchecker.WinnerResponse, error) {
	mark, err := c.checker.Winner(ctx, game.ID(req.GameId))
	if err != nil {
		return nil, err
	}
	if mark == nil {
		return &rpcchecker.WinnerResponse{
			HasWinner: false,
		}, nil
	}
	m := *mark
	return &rpcchecker.WinnerResponse{
		Mark:      string(m),
		HasWinner: true,
	}, nil
}
