package botapiclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/theothertomelliott/tic-tac-toverengineered/api/pkg/apiclient"
	"github.com/theothertomelliott/tic-tac-toverengineered/bot/pkg/bot"
	"github.com/theothertomelliott/tic-tac-toverengineered/checker/pkg/win"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
)

func New(baseURL string, bot bot.Bot) *Client {
	return &Client{
		bot: bot,
		api: apiclient.New(baseURL, http.DefaultClient),
	}
}

type Client struct {
	bot bot.Bot
	api *apiclient.Client
}

// PlayGame starts a new game and plays it to conclusion as both players.
func (c *Client) PlayGame(ctx context.Context) (string, error) {
	// Create a new game
	var gameID game.ID
	if err := c.api.RawApiGet(ctx, "/new", &gameID); err != nil {
		return "", fmt.Errorf("could not create game: %w", err)
	}

	log.Printf("%v> Starting game", gameID)

	// Alternate turns until a winner
	for true {
		if won, err := c.hasWinner(ctx, gameID); won != "" || err != nil {
			return won, err
		}

		p, err := c.currentPlayer(ctx, gameID)
		if err != nil {
			return "", err
		}

		var g [][]*player.Mark
		if err := c.api.ApiGet(ctx, gameID, "grid", &g); err != nil {
			return "", fmt.Errorf("could not load grid: %w", err)
		}

		pos, err := c.bot.Move(p, g)
		if err != nil {
			return "", err
		}
		jPos, err := json.Marshal(pos)
		if err != nil {
			return "", err
		}

		log.Printf("%v> %v plays %s", gameID, p, string(jPos))
		resp, err := c.api.Get(ctx, gameID, fmt.Sprintf("play?player=%v&pos=%v", p, string(jPos)))
		if err != nil {
			return "", fmt.Errorf("could not take turn: %w", err)
		}
		if resp.StatusCode != http.StatusOK {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			var msg string = string(body)
			if err != nil {
				msg = err.Error()
			}
			return "", fmt.Errorf("could not take turn: %v", msg)
		}

	}

	return "", nil
}

func (c *Client) hasWinner(ctx context.Context, gameID game.ID) (string, error) {
	var r win.Result
	if err := c.api.ApiGet(ctx, gameID, "winner", &r); err != nil {
		return "", fmt.Errorf("could not determine winner: %v", err)
	}
	if !r.Finished() {
		return "", nil
	}
	if r.IsDraw {
		return "Draw", nil
	}
	return r.Winner.String(), nil
}

func (c Client) currentPlayer(ctx context.Context, gameID game.ID) (player.Mark, error) {
	var m player.Mark
	if err := c.api.ApiGet(ctx, gameID, "player/current", &m); err != nil {
		return player.X, fmt.Errorf("could not determine player: %w", err)
	}
	return m, nil
}
