package main

import (
	"log"
	"net/http"
	"os"

	"github.com/theothertomelliott/tic-tac-toverengineered/internal/api"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game/rpcrepository/repoclient"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/turn/inmemoryturns"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/win/gridchecker"
)

func getRepoServerTarget() string {
	if repoServerTarget := os.Getenv("REPO_SERVER_TARGET"); repoServerTarget != "" {
		return repoServerTarget
	}
	return "localhost:8082"
}

func main() {
	log.Println("Starting api server")
	mux := http.NewServeMux()
	g := grid.NewInMemory()
	checker := gridchecker.New(g)
	controller := inmemoryturns.New(
		inmemoryturns.NewCurrentTurn(),
		g,
		checker,
	)
	r, err := repoclient.Connect(getRepoServerTarget())
	if err != nil {
		log.Fatalf("could not connect to repo server: %v", err)
	}

	server := api.New(r, controller, g, checker)
	server.CreateRoutes(mux)

	log.Println("Listening on port :8080")
	http.ListenAndServe(":8080", mux)
}
