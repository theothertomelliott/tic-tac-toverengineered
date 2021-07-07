package botopenapiclient

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/bot/pkg/bot"
)

func New(baseURL string, bot bot.Bot) (*Client, error) {
	apiClient, err := tictactoeapi.NewClientWithResponses(baseURL)
	if err != nil {
		return nil, err
	}
	return &Client{
		bot: bot,
		api: apiClient,
	}, nil
}

type Client struct {
	bot bot.Bot
	api tictactoeapi.ClientWithResponsesInterface
}

// PlayGame starts a new game and plays it to conclusion as both players.
func (c *Client) PlayGame(ctx context.Context) (string, error) {

	player1, player2, err := c.createGame(ctx)
	if err != nil {
		return "", fmt.Errorf("creating game: %w", err)
	}

	if player1.GameID != player2.GameID {
		log.Fatal("Game IDs did not match")
	}

	fmt.Println(player1, player2)

	// Alernate turns until we have a winner
	currentPlayer := player1
	for {
		winner, err := c.api.WinnerWithResponse(ctx, player1.GameID)
		if err != nil {
			return "", fmt.Errorf("checking winner: %w", err)
		}
		if winner.JSON200 == nil {
			return "", fmt.Errorf("Got response %v: %v", winner.StatusCode(), string(winner.Body))
		}
		if winner.JSON200.Winner != nil {
			return *winner.JSON200.Winner, nil
		}
		if winner.JSON200.Draw != nil && *winner.JSON200.Draw {
			return "Draw", nil
		}

		// Take the current player's turn
		err = c.takeTurn(ctx, currentPlayer)
		if err != nil {
			return "", fmt.Errorf("taking turn: %w", err)
		}

		// Switch player
		if currentPlayer == player1 {
			currentPlayer = player2
		} else {
			currentPlayer = player1
		}

	}
}

// takeTurn loads the current grid, identifies the bot's next move and then
// makes that move in the current game.
func (c *Client) takeTurn(ctx context.Context, player *tictactoeapi.Match) error {
	gridResponse, err := c.api.GameGridWithResponse(ctx, player.GameID)
	if err != nil {
		return err
	}
	if gridResponse.JSON200 == nil {
		return fmt.Errorf("Got response %v: %v", gridResponse.StatusCode(), string(gridResponse.Body))
	}
	grid := gridResponse.JSON200
	pos, err := move(player.Mark, grid.Grid)
	if err != nil {
		return err
	}

	playRes, err := c.api.PlayWithResponse(ctx, player.GameID, &tictactoeapi.PlayParams{
		Position: pos,
		Token:    player.Token,
	})
	if err != nil {
		return err
	}
	if playRes.JSON200 == nil {
		return fmt.Errorf("Got response %v: %v", playRes.StatusCode(), string(playRes.Body))
	}
	return nil
}

func move(mark string, state [][]string) (tictactoeapi.Position, error) {
	var valid []tictactoeapi.Position
	for i, row := range state {
		for j, m := range row {
			if m == "" {
				valid = append(
					valid,
					tictactoeapi.Position{
						I: int32(i),
						J: int32(j),
					},
				)
			}
		}
	}
	if len(valid) == 0 {
		return tictactoeapi.Position{}, fmt.Errorf("no valid moves")
	}
	return valid[rand.Intn(len(valid))], nil
}

func (c *Client) createGame(ctx context.Context) (player1 *tictactoeapi.Match, player2 *tictactoeapi.Match, err error) {
	mp1, err := c.getMatchPending(ctx)
	if err != nil {
		return nil, nil, err
	}
	mp2, err := c.getMatchPending(ctx)
	if err != nil {
		return nil, nil, err
	}

	player1, err = c.getMatch(ctx, mp1)
	if err != nil {
		return nil, nil, err
	}

	player2, err = c.getMatch(ctx, mp2)
	if err != nil {
		return nil, nil, err
	}

	return player1, player2, nil
}

func (c *Client) getMatchPending(ctx context.Context) (*tictactoeapi.MatchPending, error) {
	matchPending, err := c.api.RequestMatchWithResponse(ctx)
	if err != nil {
		return nil, err
	}
	if matchPending.JSON202 == nil {
		return nil, fmt.Errorf("Got response %v: %v", matchPending.StatusCode(), string(matchPending.Body))
	}
	return matchPending.JSON202, nil
}

func (c *Client) getMatch(ctx context.Context, matchPending *tictactoeapi.MatchPending) (*tictactoeapi.Match, error) {
	matchResult, err := c.api.MatchStatusWithResponse(ctx, &tictactoeapi.MatchStatusParams{
		RequestID: matchPending.RequestID,
	})
	if err != nil {
		return nil, err
	}
	if matchResult.JSON200 == nil {
		return nil, fmt.Errorf("Got response %v: %v", matchResult.StatusCode(), string(matchResult.Body))
	}
	return matchResult.JSON200, nil
}
