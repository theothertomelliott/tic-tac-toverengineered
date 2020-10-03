package main

import (
	"log"
	"net/http"
	"os"

	api "github.com/theothertomelliott/tic-tac-toverengineered/api/internal"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game/rpcrepository/repoclient"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/grid/rpcgrid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/turn/rpcturn"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/win/rpcchecker"
)

func getCurrentTurnServerTarget() string {
	if serverTarget := os.Getenv("CURRENT_TURN_SERVER_TARGET"); serverTarget != "" {
		return serverTarget
	}
	return "localhost:8084"
}

func getRepoServerTarget() string {
	if serverTarget := os.Getenv("REPO_SERVER_TARGET"); serverTarget != "" {
		return serverTarget
	}
	return "localhost:8082"
}

func getGridServerTarget() string {
	if serverTarget := os.Getenv("GRID_SERVER_TARGET"); serverTarget != "" {
		return serverTarget
	}
	return "localhost:8086"
}

func getCheckerServerTarget() string {
	if serverTarget := os.Getenv("CHECKER_SERVER_TARGET"); serverTarget != "" {
		return serverTarget
	}
	return "localhost:8088"
}

func getTurnControllerServerTarget() string {
	if serverTarget := os.Getenv("TURN_CONTROLLER_SERVER_TARGET"); serverTarget != "" {
		return serverTarget
	}
	return "localhost:8090"
}

func main() {
	log.Println("Starting api server")
	mux := http.NewServeMux()
	g, err := rpcgrid.ConnectGrid(getGridServerTarget())
	if err != nil {
		log.Fatalf("could not connect to grid server: %v", err)
	}
	checker, err := rpcchecker.ConnectChecker(getCheckerServerTarget())
	if err != nil {
		log.Fatalf("could not connect to win checker server: %v", err)
	}
	controller, err := rpcturn.ConnectController(getTurnControllerServerTarget())
	if err != nil {
		log.Fatalf("could not connect to turn controller server: %v", err)
	}
	r, err := repoclient.Connect(getRepoServerTarget())
	if err != nil {
		log.Fatalf("could not connect to repo server: %v", err)
	}

	server := api.New(r, controller, g, checker)
	server.CreateRoutes(mux)

	log.Println("Listening on port :8080")
	http.ListenAndServe(":8080", mux)
}
