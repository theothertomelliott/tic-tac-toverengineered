package botopenapiclient

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi/tictactoeapiclient"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/bot/pkg/bot"
)

func New(baseURL string, bot bot.Bot) (*Client, error) {
	apiClient, err := tictactoeapi.NewClientWithResponses(baseURL)
	if err != nil {
		return nil, err
	}
	return &Client{
		bot: bot,
		api: tictactoeapiclient.New(apiClient),
	}, nil
}

type Client struct {
	bot bot.Bot
	api *tictactoeapiclient.Client
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
		winner, err := c.api.Winner(ctx, currentPlayer.GameID)
		if err != nil {
			return "", fmt.Errorf("checking winner: %w", err)
		}
		if winner.Winner != nil {
			return *winner.Winner, nil
		}
		if winner.Draw != nil && *winner.Draw {
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
	grid, err := c.api.GameGrid(ctx, player.GameID)
	if err != nil {
		return err
	}
	pos, err := move(player.Mark, grid)
	if err != nil {
		return err
	}

	return c.api.Play(ctx, player, pos)
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

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
	return valid[r.Intn(len(valid))], nil
}

func (c *Client) createGame(ctx context.Context) (player1 *tictactoeapi.Match, player2 *tictactoeapi.Match, err error) {
	requestID1, err := c.api.RequestMatch(ctx)
	if err != nil {
		return nil, nil, err
	}
	requestID2, err := c.api.RequestMatch(ctx)
	if err != nil {
		return nil, nil, err
	}

	player1, err = c.api.MatchStatus(ctx, requestID1)
	if err != nil {
		return nil, nil, err
	}

	player2, err = c.api.MatchStatus(ctx, requestID2)
	if err != nil {
		return nil, nil, err
	}

	return player1, player2, nil
}
