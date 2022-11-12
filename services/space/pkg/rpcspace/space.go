package rpcspace

import (
	context "context"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/filters"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func ConnectSpace(target string) (*Space, error) {
	var err error
	c := &Space{}
	c.conn, err = grpc.Dial(
		target,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor(
			// Do not trace health checks
			otelgrpc.WithInterceptorFilter(
				filters.Not(
					filters.HealthCheck(),
				),
			))),
	)
	if err != nil {
		return nil, err
	}
	c.client = NewSpaceClient(c.conn)
	c.health = grpc_health_v1.NewHealthClient(c.conn)
	return c, nil
}

type Space struct {
	conn   *grpc.ClientConn
	client SpaceClient
	health grpc_health_v1.HealthClient
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

func (s *Space) Health() grpc_health_v1.HealthClient {
	return s.health
}
