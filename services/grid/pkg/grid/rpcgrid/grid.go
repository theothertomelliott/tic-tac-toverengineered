package rpcgrid

import (
	context "context"
	"log"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid"
	grpc "google.golang.org/grpc"
)

func ConnectGrid(target string) (*Grid, error) {
	var err error
	c := &Grid{}
	c.conn, err = grpc.Dial(
		target,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(monitoring.UnaryClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}
	c.client = NewGridClient(c.conn)
	return c, nil

}

type Grid struct {
	conn   *grpc.ClientConn
	client GridClient
}

func PositionToProtoPosition(pos grid.Position) *Position {
	return &Position{
		X: int32(pos.X),
		Y: int32(pos.Y),
	}
}

func ProtoPositionToPosition(pos *Position) grid.Position {
	return grid.Position{
		X: int(pos.X),
		Y: int(pos.Y),
	}
}

func (g *Grid) Mark(ctx context.Context, id game.ID, pos grid.Position) (*player.Mark, error) {
	resp, err := g.client.Mark(ctx, &MarkRequest{
		GameId:   string(id),
		Position: PositionToProtoPosition(pos),
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

func (g *Grid) State(ctx context.Context, id game.ID) ([][]*player.Mark, error) {
	gameState, err := g.client.State(ctx, &StateRequest{
		GameId: string(id),
	})
	if err != nil {
		return nil, err
	}
	var out [][]*player.Mark
	for _, rowState := range gameState.RowState {
		var row []*player.Mark
		for _, m := range rowState.Mark {
			var mark *player.Mark
			if m.HasMark {
				markVal := player.Mark(m.Mark)
				mark = &markVal
			}
			row = append(row, mark)
		}
		out = append(out, row)
	}
	return out, nil
}

func (g *Grid) SetMark(ctx context.Context, id game.ID, pos grid.Position, m player.Mark) error {
	_, err := g.client.SetMark(ctx, &SetMarkRequest{
		GameId:   string(id),
		Position: PositionToProtoPosition(pos),
		Mark:     string(m),
	})
	return err
}

func (g *Grid) Rows(ctx context.Context) []grid.Row {
	resp, err := g.client.Rows(ctx, &RowsRequest{})
	if err != nil {
		log.Println("Error getting rows:", err)
		panic(err)
	}
	var out []grid.Row
	for _, row := range resp.Row {
		var r grid.Row
		for _, p := range row.Position {
			r = append(r, ProtoPositionToPosition(p))
		}
		out = append(out, r)
	}
	return out
}
