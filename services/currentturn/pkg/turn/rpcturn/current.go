package rpcturn

import (
	context "context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	grpc "google.golang.org/grpc"
)

func ConnectCurrent(target string) (*Current, error) {
	var err error
	c := &Current{}
	c.conn, err = grpc.Dial(
		target,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(monitoring.UnaryClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}
	c.client = NewCurrentClient(c.conn)
	return c, nil

}

type Current struct {
	conn   *grpc.ClientConn
	client CurrentClient
}

func (c *Current) Close() error {
	return c.conn.Close()
}

func (c *Current) Player(ctx context.Context, id game.ID) (player.Mark, error) {
	resp, err := c.client.Player(ctx, &PlayerRequest{GameId: string(id)})
	if err != nil {
		return player.Mark(""), err
	}
	return player.Mark(resp.Mark), nil
}

func (c *Current) Next(ctx context.Context, id game.ID) error {
	_, err := c.client.Next(ctx, &NextRequest{GameId: string(id)})
	if err != nil {
		return err
	}
	return nil
}
