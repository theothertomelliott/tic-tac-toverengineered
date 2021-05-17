package grid

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid/rpcgrid"
)

func NewServer(grid grid.Grid) rpcgrid.GridServer {
	return &gridServer{
		grid: grid,
	}
}

type gridServer struct {
	rpcgrid.UnimplementedGridServer

	grid grid.Grid
}

func (g *gridServer) Mark(ctx context.Context, req *rpcgrid.MarkRequest) (*rpcgrid.MarkResponse, error) {
	mark, err := g.grid.Mark(ctx, game.ID(req.GameId), rpcgrid.ProtoPositionToPosition(req.Position))
	if err != nil {
		return nil, err
	}
	if mark == nil {
		return &rpcgrid.MarkResponse{
			HasMark: false,
		}, nil
	}
	m := *mark
	return &rpcgrid.MarkResponse{
		Mark:    string(m),
		HasMark: true,
	}, nil
}

func (g *gridServer) State(ctx context.Context, req *rpcgrid.StateRequest) (*rpcgrid.StateResponse, error) {
	state, err := g.grid.State(ctx, game.ID(req.GameId))
	if err != nil {
		return nil, err
	}
	resp := &rpcgrid.StateResponse{}
	for _, row := range state {
		rs := &rpcgrid.RowState{}
		for _, m := range row {
			mr := &rpcgrid.MarkResponse{}
			if m != nil {
				mr.HasMark = true
				mr.Mark = string(*m)
			}
			rs.Mark = append(rs.Mark, mr)
		}
		resp.RowState = append(resp.RowState, rs)
	}
	return resp, nil
}

func (g *gridServer) SetMark(ctx context.Context, req *rpcgrid.SetMarkRequest) (*rpcgrid.SetMarkResponse, error) {
	return &rpcgrid.SetMarkResponse{},
		g.grid.SetMark(
			ctx,
			game.ID(req.GameId),
			rpcgrid.ProtoPositionToPosition(req.Position),
			player.Mark(req.Mark),
		)
}

func (g *gridServer) Rows(ctx context.Context, req *rpcgrid.RowsRequest) (*rpcgrid.RowsResponse, error) {
	resp := &rpcgrid.RowsResponse{}
	rows := g.grid.Rows(ctx)
	for _, row := range rows {
		var r rpcgrid.Row
		for _, p := range row {
			r.Position = append(r.Position, rpcgrid.PositionToProtoPosition(p))
		}
		resp.Row = append(resp.Row, &r)
	}
	return resp, nil
}
