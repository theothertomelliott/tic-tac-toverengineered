package tictactoeapiclient

import (
	"context"
	"fmt"

	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
)

func New(client tictactoeapi.ClientWithResponsesInterface) *Client {
	return &Client{
		client: client,
	}
}

// Client is a wrapper around a tictactoeapi client that provides convenience methods.
type Client struct {
	client tictactoeapi.ClientWithResponsesInterface
}

// Index retrieves a list of games
func (c *Client) Index(ctx context.Context, offset *int64, max *int64) ([]string, error) {
	res, err := c.client.IndexWithResponse(ctx, &tictactoeapi.IndexParams{
		Offset: offset,
		Max:    max,
	})
	if err != nil {
		return nil, err
	}

	if res.JSONDefault != nil {
		return nil, fmt.Errorf(res.JSONDefault.Message)
	}
	if res.JSON200 == nil {
		return nil, fmt.Errorf("unexpected response %q: %v", res.Status(), string(res.Body))
	}
	return *res.JSON200, nil
}

// RequestMatch makes a request for a new game and returns the request ID for status updates.
func (c *Client) RequestMatch(ctx context.Context) (string, error) {
	res, err := c.client.RequestMatchWithResponse(ctx)
	if err != nil {
		return "", err
	}
	if res.JSONDefault != nil {
		return "", fmt.Errorf(res.JSONDefault.Message)
	}
	if res.JSON202 == nil {
		return "", fmt.Errorf("unexpected response %q: %v", res.Status(), string(res.Body))
	}
	return res.JSON202.RequestID, nil
}

// MatchStatus retrieves the status of a match request.
// The Match is returned if a match has been made, otherwise nil.
func (c *Client) MatchStatus(ctx context.Context, requestID string) (*tictactoeapi.Match, error) {
	res, err := c.client.MatchStatusWithResponse(ctx, &tictactoeapi.MatchStatusParams{
		RequestID: requestID,
	})
	if err != nil {
		return nil, err
	}
	if res.JSONDefault != nil {
		return nil, fmt.Errorf(res.JSONDefault.Message)
	}
	if res.JSON200 == nil {
		return nil, fmt.Errorf("unexpected response %q: %v", res.Status(), string(res.Body))
	}
	if res.JSON102 != nil {
		return nil, nil
	}
	return res.JSON200, nil
}

// GameGrid retrieves the current state of a game.
func (c *Client) GameGrid(ctx context.Context, gameID string) ([][]string, error) {
	res, err := c.client.GameGridWithResponse(ctx, gameID)
	if err != nil {
		return nil, err
	}
	if res.JSONDefault != nil {
		return nil, fmt.Errorf(res.JSONDefault.Message)
	}
	if res.JSON200 == nil {
		return nil, fmt.Errorf("unexpected response %q: %v", res.Status(), string(res.Body))
	}
	return res.JSON200.Grid, nil
}

// Play makes a move in a game.
func (c *Client) Play(ctx context.Context, match *tictactoeapi.Match, pos tictactoeapi.Position) error {
	playRes, err := c.client.PlayWithResponse(ctx, match.GameID, &tictactoeapi.PlayParams{
		Token:    match.Token,
		Position: pos,
	})
	if err != nil {
		return err
	}
	if playRes.JSONDefault != nil {
		return fmt.Errorf(playRes.JSONDefault.Message)
	}
	if playRes.JSON200 == nil {
		return fmt.Errorf("unexpected response %q: %v", playRes.Status(), string(playRes.Body))
	}
	return nil
}

// CurrentPlayer returns the player in a game who should next make a move.
func (c *Client) CurrentPlayer(ctx context.Context, gameID string) (string, error) {
	res, err := c.client.CurrentPlayerWithResponse(ctx, gameID)
	if err != nil {
		return "", err
	}
	if res.JSONDefault != nil {
		return "", fmt.Errorf(res.JSONDefault.Message)
	}
	if res.JSON200 == nil {
		return "", fmt.Errorf("unexpected response %q: %v", res.Status(), string(res.Body))
	}
	return *res.JSON200, nil
}

// Winner returns the winner of a game, if any.
func (c *Client) Winner(ctx context.Context, gameID string) (tictactoeapi.Winner, error) {
	res, err := c.client.WinnerWithResponse(ctx, gameID)
	if err != nil {
		return tictactoeapi.Winner{}, err
	}
	if res.JSONDefault != nil {
		return tictactoeapi.Winner{}, fmt.Errorf(res.JSONDefault.Message)
	}
	if res.JSON200 == nil {
		return tictactoeapi.Winner{}, fmt.Errorf("unexpected response %q: %v", res.Status(), string(res.Body))
	}
	return *res.JSON200, nil
}
