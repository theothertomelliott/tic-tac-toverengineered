package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/env"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/bot/pkg/bot/randombot"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/bot/pkg/botopenapiclient"
)

var (
	gamesProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "bot_games_processed",
		Help: "The total number of bots that played a game",
	})
	gamesFailed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "bot_games_failed",
		Help: "The total number of games that failed for a bot",
	})
)

func main() {
	apiBaseURL := env.Get("API_BASE_URL", "http://localhost:8094")

	c, err := botopenapiclient.New(apiBaseURL, randombot.New())
	if err != nil {
		log.Fatal(err)
	}

	err = c.WaitForReady(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var results = make(chan bool)

	go player("Player 1", c, results)
	go player("Player 2", c, results)
	go player("Player 3", c, results)

	go func() {
		for result := range results {
			gamesProcessed.Inc()
			if !result {
				gamesFailed.Inc()
			}
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

func player(name string, c *botopenapiclient.Client, results chan<- bool) {
	for {
		time.Sleep(time.Second / 4)
		win, err := c.Play(context.Background(), name)
		if err != nil {
			log.Printf("Failed game - %v: %v", name, err)
			results <- false
			continue
		}
		log.Printf("%v won? %v\n", name, win)
		results <- true
	}
}
