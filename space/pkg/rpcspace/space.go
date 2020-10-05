package rpcspace

import (
	context "context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
	grpc "google.golang.org/grpc"
)

func ConnectSpace(target string) (*Space, error) {
	var err error
	c := &Space{}
	c.conn, err = grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c.client = NewSpaceClient(c.conn)
	return c, nil
}

type Space struct {
	conn   *grpc.ClientConn
	client SpaceClient
}

func (c *Space) Close() error {
	return c.conn.Close()
}

func (s *Space) Mark(ctx context.Context, id game.ID) (*player.Mark, error) {
	resp, err := s.client.Mark(ctx, &MarkRequest{
		GameId: string(id),
	})
	if err != nil {
		return nil, err
	}
	if !resp.HasMark {
		return nil, nil
	}
	m := player.Mark(resp.Mark)
	return &m, nil
}

func (s *Space) SetMark(ctx context.Context, id game.ID, m player.Mark) error {
	_, err := s.client.SetMark(ctx, &SetMarkRequest{
		GameId: string(id),
		Mark:   string(m),
	})
	return err
}
