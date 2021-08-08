package main

import (
	"context"
	"log"
	"time"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/env"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/bot/pkg/bot/randombot"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/bot/pkg/botopenapiclient"
)

func main() {
	apiBaseURL := env.Get("API_BASE_URL", "http://localhost:8094")

	c, err := botopenapiclient.New(apiBaseURL, randombot.New())
	if err != nil {
		log.Fatal(err)
	}

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
		time.Sleep(3 * time.Minute)
	}

}
