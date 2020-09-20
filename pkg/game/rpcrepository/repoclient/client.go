package repoclient

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game/rpcrepository"
	"google.golang.org/grpc"
)

// Connect establishes a connection to a repository server and returns a
// client.
func Connect(target string) (*Client, error) {
	var err error
	c := &Client{}
	c.conn, err = grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c.client = rpcrepository.NewRepoClient(c.conn)
	return c, nil
}

type Client struct {
	conn   *grpc.ClientConn
	client rpcrepository.RepoClient
}

func (c *Client) Close() error {
	return c.conn.Close()
}

// New creates a new game and creates a unique ID
func (c *Client) New(ctx context.Context) (game.ID, error) {
	resp, err := c.client.New(ctx, &rpcrepository.NewRequest{})
	if err != nil {
		return game.ID(""), err
	}
	return game.ID(resp.ID), nil
}

// List obtains game IDs, ordered by creation date.
// Pagination is managed through the max and offset params.
func (c *Client) List(ctx context.Context, max int64, offset int64) ([]game.ID, error) {
	resp, err := c.client.List(ctx, &rpcrepository.ListRequest{Max: max, Offset: offset})
	if err != nil {
		return nil, err
	}
	var out []game.ID
	for _, id := range resp.ID {
		out = append(out, game.ID(id))
	}
	return out, nil
}

// Exists returns true iff the given game ID was previously created with New
func (c *Client) Exists(ctx context.Context, id game.ID) (bool, error) {
	resp, err := c.client.Exists(ctx, &rpcrepository.ExistsRequest{ID: string(id)})
	if err != nil {
		return false, err
	}
	return resp.Exists, nil
}
