package botopenapiclient

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
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

// WaitForReady blocks until the api server is ready
func (c *Client) WaitForReady(ctx context.Context) error {
	for true {
		_, err := c.api.Index(ctx, nil, nil)
		if err == nil {
			log.Printf("api is ready")
			return nil
		}
		log.Printf("waiting for ready: %v", err)
		time.Sleep(time.Second)
	}
	return nil
}

const (
	matchTimeoutDuration = time.Minute
	turnTimeoutDuration  = time.Second * 5
)

// Play joins a game, plays it synchronously and returns whether or not this player won.
func (c *Client) Play(ctx context.Context, name string) (bool, error) {
	// Request a game
	log.Printf("%v: Requesting match", name)
	requestID, err := c.api.RequestMatch(ctx)
	if err != nil {
		return false, fmt.Errorf("requesting match: %v", err)
	}

	// Wait until a match has been made
	var match *tictactoeapi.Match
	timeout := time.After(matchTimeoutDuration)
	for match == nil {
		select {
		case <-timeout:
			return false, fmt.Errorf("timed out waiting for match")
		default:
		}
		match, err = c.api.MatchStatus(ctx, requestID)
		if err != nil {
			return false, fmt.Errorf("checking match status: %v", err)
		}
		if match == nil {
			time.Sleep(time.Millisecond)
		}
	}

	log.Printf("%v: match = %+v", name, match)

	// Play until the game has been won

	var winner tictactoeapi.Winner
	var lastMove = time.Now()
	for {
		if time.Since(lastMove) > turnTimeoutDuration {
			return false, fmt.Errorf("timed out waiting for opponent (%v)", match.GameID)
		}
		time.Sleep(time.Millisecond)
		// Check winner
		winner, err = c.api.Winner(ctx, match.GameID)
		if err != nil {
			return false, fmt.Errorf("checking winner (%v): %v", match.GameID, err)
		}
		if winner.Winner != nil {
			return *winner.Winner == match.Mark, nil
		}
		if winner.Draw != nil && *winner.Draw {
			return false, nil
		}

		// Check current player
		currentPlayer, err := c.api.CurrentPlayer(ctx, match.GameID)
		if err != nil {
			return false, fmt.Errorf("getting current player (%v): %v", match.GameID, err)
		}
		if currentPlayer != match.Mark {
			continue
		}

		grid, err := c.api.GameGrid(ctx, match.GameID)
		if err != nil {
			return false, fmt.Errorf("getting game grid (%v): %v", match.GameID, err)
		}
		pos, err := c.bot.Move(player.Mark(match.Mark), grid)
		if err != nil {
			return false, fmt.Errorf("identifying move (%v): %v", match.GameID, err)
		}

		// Check winner again
		winner, err = c.api.Winner(ctx, match.GameID)
		if err != nil {
			return false, fmt.Errorf("checking winner (%v): %v", match.GameID, err)
		}
		if winner.Winner != nil {
			return *winner.Winner == match.Mark, nil
		}
		if winner.Draw != nil && *winner.Draw {
			return false, nil
		}

		log.Printf("%v %v %v: playing %v", name, match.GameID, match.Mark, pos)
		err = c.api.Play(ctx, match.GameID, match.Token, pos)
		if err != nil {
			return false, fmt.Errorf("making move (%v): %v", match.GameID, err)
		}
		lastMove = time.Now()
	}

	return false, nil
}

// PlayBothSides starts a new game and plays it to conclusion as both players.
func (c *Client) PlayBothSides(ctx context.Context) (string, error) {
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
func (c *Client) takeTurn(ctx context.Context, p *tictactoeapi.Match) error {
	grid, err := c.api.GameGrid(ctx, p.GameID)
	if err != nil {
		return err
	}
	pos, err := c.bot.Move(player.Mark(p.Mark), grid)
	if err != nil {
		return err
	}

	return c.api.Play(ctx, p.GameID, p.Token, pos)
}

func (c *Client) createGame(ctx context.Context) (player1 *tictactoeapi.Match, player2 *tictactoeapi.Match, err error) {
	players, err := c.api.RequestMatchPair(ctx)
	if err != nil {
		return nil, nil, err
	}

	return &players.X, &players.O, nil
}
