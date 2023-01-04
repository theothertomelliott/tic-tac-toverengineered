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

func (g *gridServer) Mark(ctx context.Context, req *rpcgrid.MarkAtPositionRequest) (*rpcgrid.MarkAtPositionResponse, error) {
	mark, err := g.grid.Mark(ctx, game.ID(req.GameId), rpcgrid.ProtoPositionToPosition(req.Position))
	if err != nil {
		return nil, err
	}
	if mark == nil {
		return &rpcgrid.MarkAtPositionResponse{
			HasMark: false,
		}, nil
	}
	m := *mark
	return &rpcgrid.MarkAtPositionResponse{
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
			mr := &rpcgrid.MarkAtPositionResponse{}
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

func (g *gridServer) SetMarkAtPosition(ctx context.Context, req *rpcgrid.SetMarkAtPositionRequest) (*rpcgrid.SetMarkAtPositionResponse, error) {
	return &rpcgrid.SetMarkAtPositionResponse{},
		g.grid.SetMark(
			ctx,
			game.ID(req.GameId),
			rpcgrid.ProtoPositionToPosition(req.Position),
			player.Mark(req.Mark),
		)
}

func (g *gridServer) Rows(ctx context.Context, req *rpcgrid.RowsRequest) (*rpcgrid.RowsResponse, error) {
	resp := &rpcgrid.RowsResponse{}
	rows, err := g.grid.Rows(ctx)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		var r rpcgrid.Row
		for _, p := range row {
			r.Position = append(r.Position, rpcgrid.PositionToProtoPosition(p))
		}
		resp.Row = append(resp.Row, &r)
	}
	return resp, nil
}
