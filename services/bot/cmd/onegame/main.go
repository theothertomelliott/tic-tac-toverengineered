package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/bot/pkg/bot/randombot"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/bot/pkg/botapiclient"
)

func main() {
	version.Println()

	apiURL := os.Args[1]
	fmt.Printf("Will use server: %v\n", apiURL)
	c := botapiclient.New(apiURL, randombot.New())
	w, err := c.PlayGame(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Winner: %v\n", w)
}
