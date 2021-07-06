package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/env"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/server"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/checker/pkg/win/rpcchecker"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/currentturn/pkg/turn/rpcturn"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game/rpcrepository/repoclient"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid/rpcgrid"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/matchmaker/rpcmatchmaker"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/matchmaker/unsignedtokens"
)

func main() {
	log.Println("Starting api server")
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
	m, err := rpcmatchmaker.Connect(getMatchMakerServerTarget())
	if err != nil {
		log.Fatalf("could not connect to matchmaker server: %v", err)
	}

	apiServer := server.New(
		r,
		m,
		controller,
		g,
		checker,
		&unsignedtokens.UnsignedTokens{},
	)

	e := echo.New()
	tictactoeapi.RegisterHandlers(e, apiServer)

	port := env.Get("PORT", "8080")
	err = e.Start(fmt.Sprintf(":%v", port))
	if err != nil {
		panic(err)
	}
}

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

func getMatchMakerServerTarget() string {
	if serverTarget := os.Getenv("MATCHMAKER_SERVER_TARGET"); serverTarget != "" {
		return serverTarget
	}
	return "localhost:8092"
}
