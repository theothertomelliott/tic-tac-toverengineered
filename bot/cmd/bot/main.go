package main

import (
	"context"
	"log"
	"time"

	"github.com/theothertomelliott/tic-tac-toverengineered/bot/pkg/bot/randombot"
	"github.com/theothertomelliott/tic-tac-toverengineered/bot/pkg/botapiclient"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/env"
)

func main() {
	apiBaseURL := env.Get("API_BASE_URL", "http://localhost:8081")

	c := botapiclient.New(apiBaseURL, randombot.New())

	// Repeatedly create games
	for true {
		winner, err := c.PlayGame(context.Background())
		if err != nil {
			log.Printf("Failed game: %q. Sleeping for 10s then resuming.", err)
			time.Sleep(10 * time.Second)
			continue
		}
		log.Printf("Winner: %v", winner)
		// Pause to slow down the rate of events
		time.Sleep(5 * time.Minute)
	}

}
