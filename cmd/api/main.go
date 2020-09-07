package main

import (
	"net/http"

	"github.com/theothertomelliott/tic-tac-toverengineered/internal/api"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/turn/inmemoryturns"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/win/gridchecker"
)

func main() {
	mux := http.NewServeMux()
	g := grid.NewInMemory()
	checker := gridchecker.New(g)
	controller := inmemoryturns.New(
		inmemoryturns.NewCurrentTurn(),
		g,
		checker,
	)

	server := api.New(controller, g, checker)
	server.CreateRoutes(mux)

	http.ListenAndServe(":8081", mux)
}
