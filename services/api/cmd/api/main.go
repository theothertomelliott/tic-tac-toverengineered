package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/defaultmonitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	api "github.com/theothertomelliott/tic-tac-toverengineered/services/api/internal"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/checker/pkg/win/rpcchecker"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/currentturn/pkg/turn/rpcturn"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game/rpcrepository/repoclient"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid/rpcgrid"
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
	version.Println()
	defaultmonitoring.Init("api")
	defer monitoring.Close()

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

	m := mux.NewRouter()
	server := api.New(r, controller, g, checker)

	prefix := os.Getenv("ROUTE_PREFIX")
	if prefix != "" && prefix != "/" {
		m = m.PathPrefix(prefix).Subrouter()
	}
	server.AddRoutes(m)

	log.Println("Listening on port :8080")
	http.ListenAndServe(":8080", m)
}

type wrappedHandler struct {
	h http.Handler
}

func (p *wrappedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Serving")
	fmt.Fprintln(w, r.URL.Path)
	p.h.ServeHTTP(w, r)
}
