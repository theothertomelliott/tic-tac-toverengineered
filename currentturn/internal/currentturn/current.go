package currentturn

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/currentturn/pkg/turn"
	"github.com/theothertomelliott/tic-tac-toverengineered/currentturn/pkg/turn/rpcturn"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
)

func NewServer(current turn.Current) rpcturn.CurrentServer {
	return &currentServer{
		current: current,
	}
}

type currentServer struct {
	current turn.Current
}

func (c *currentServer) Player(ctx context.Context, req *rpcturn.PlayerRequest) (*rpcturn.PlayerResponse, error) {
	resp := &rpcturn.PlayerResponse{}
	mark, err := c.current.Player(ctx, game.ID(req.GameId))
	if err != nil {
		return nil, err
	}
	resp.Mark = string(mark)
	return resp, nil
}

func (c *currentServer) Next(ctx context.Context, req *rpcturn.NextRequest) (*rpcturn.NextResponse, error) {
	resp := &rpcturn.NextResponse{}
	err := c.current.Next(ctx, game.ID(req.GameId))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
