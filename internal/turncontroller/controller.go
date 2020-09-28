package turncontroller

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/turn"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/turn/rpcturn"
)

func NewServer(c turn.Controller) rpcturn.ControllerServer {
	return &controllerServer{
		controller: c,
	}
}

type controllerServer struct {
	controller turn.Controller
}

func (c *controllerServer) TakeTurn(ctx context.Context, req *rpcturn.TakeTurnRequest) (*rpcturn.TakeTurnResponse, error) {
	err := c.controller.TakeTurn(ctx, game.ID(req.GameId), player.Mark(req.Mark), rpcturn.ProtoPositionToPosition(req.Position))
	return &rpcturn.TakeTurnResponse{}, err
}

func (c *controllerServer) NextPlayer(ctx context.Context, req *rpcturn.NextPlayerRequest) (*rpcturn.NextPlayerResponse, error) {
	mark, err := c.controller.NextPlayer(ctx, game.ID(req.GameId))
	if err != nil {
		return nil, err
	}
	return &rpcturn.NextPlayerResponse{
		Mark: string(mark),
	}, nil
}
