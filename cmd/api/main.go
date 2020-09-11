package main

import (
	"log"
	"net/http"

	"github.com/theothertomelliott/tic-tac-toverengineered/internal/api"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game/inmemoryrepository"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/turn/inmemoryturns"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/win/gridchecker"
)

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
	r := inmemoryrepository.New()

	server := api.New(r, controller, g, checker)
	server.CreateRoutes(mux)

	log.Println("Listening on port :8080")
	http.ListenAndServe(":8080", mux)
}
