package space

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
	space "github.com/theothertomelliott/tic-tac-toverengineered/space/pkg"
	"github.com/theothertomelliott/tic-tac-toverengineered/space/pkg/rpcspace"
)

func NewServer(s space.Space) rpcspace.SpaceServer {
	return &spaceServer{
		space: s,
	}
}

type spaceServer struct {
	space space.Space
}

func (s *spaceServer) Mark(ctx context.Context, req *rpcspace.MarkRequest) (*rpcspace.MarkResponse, error) {
	m, err := s.space.Mark(ctx, game.ID(req.GameId))
	if err != nil {
		return nil, err
	}
	if m == nil {
		return &rpcspace.MarkResponse{
			HasMark: false,
		}, nil
	}
	return &rpcspace.MarkResponse{
		Mark:    string(*m),
		HasMark: true,
	}, nil
}

func (s *spaceServer) SetMark(ctx context.Context, req *rpcspace.SetMarkRequest) (*rpcspace.SetMarkResponse, error) {
	err := s.space.SetMark(ctx, game.ID(req.GameId), player.Mark(req.Mark))
	return &rpcspace.SetMarkResponse{}, err
}
