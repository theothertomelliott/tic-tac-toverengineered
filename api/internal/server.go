package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/checker/pkg/win"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/currentturn/pkg/turn"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid"
)

func New(
	repo game.Repository,
	turn turn.Controller,
	grid grid.Grid,
	checker win.Checker,
) *Server {
	return &Server{
		repo:    repo,
		turn:    turn,
		grid:    grid,
		checker: checker,
	}
}

type Server struct {
	repo    game.Repository
	turn    turn.Controller
	grid    grid.Grid
	checker win.Checker
}

// wrap will wrap an http handler with all intended middleware
func wrap(handler http.HandlerFunc, name string) http.Handler {
	return monitoring.WrapHTTP(handler, name)
}

func (s *Server) AddRoutes(r *mux.Router) {
	r.Handle("/", wrap(s.listGamesHandler, "index"))
	r.Handle("/{game}/grid", wrap(s.gridHandler, "grid"))
	r.Handle("/{game}/player/current", wrap(s.currentPlayerHandler, "currentplayer"))
	r.Handle("/{game}/winner", wrap(s.winnerHandler, "winner"))
	r.Handle("/{game}/play", wrap(s.playHandler, "play"))
	r.Handle("/new", wrap(s.newGameHandler, "newgame"))
}

func jsonResponse(w http.ResponseWriter, value interface{}) {
	out, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(out))
}
