package rpcturn

import (
	context "context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	grpc "google.golang.org/grpc"
)

func ConnectController(target string) (*Controller, error) {
	var err error
	c := &Controller{}
	c.conn, err = grpc.Dial(
		target,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}
	c.client = NewControllerClient(c.conn)
	return c, nil
}

type Controller struct {
	conn   *grpc.ClientConn
	client ControllerClient
}

func (c *Controller) Close() error {
	return c.conn.Close()
}

func (c *Controller) TakeTurn(ctx context.Context, id game.ID, m player.Mark, pos grid.Position) error {
	_, err := c.client.TakeTurn(ctx, &TakeTurnRequest{
		GameId:   string(id),
		Mark:     string(m),
		Position: PositionToProtoPosition(pos),
	})
	return err
}

func (c *Controller) NextPlayer(ctx context.Context, id game.ID) (player.Mark, error) {
	resp, err := c.client.NextPlayer(ctx, &NextPlayerRequest{
		GameId: string(id),
	})
	if err != nil {
		return player.Mark(""), err
	}
	return player.Mark(resp.Mark), nil
}

func PositionToProtoPosition(pos grid.Position) *TCPosition {
	return &TCPosition{
		X: int32(pos.X),
		Y: int32(pos.Y),
	}
}

func ProtoPositionToPosition(pos *TCPosition) grid.Position {
	return grid.Position{
		X: int(pos.X),
		Y: int(pos.Y),
	}
}
