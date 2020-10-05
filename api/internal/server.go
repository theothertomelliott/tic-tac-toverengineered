package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/checker/pkg/win"
	"github.com/theothertomelliott/tic-tac-toverengineered/currentturn/pkg/turn"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
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

func (s *Server) CreateRoutes(m *http.ServeMux) {
	r := mux.NewRouter()
	r.HandleFunc("/", s.listGamesHandler)
	r.HandleFunc("/{game}/grid", s.gridHandler)
	r.HandleFunc("/{game}/player/current", s.currentPlayerHandler)
	r.HandleFunc("/{game}/winner", s.winnerHandler)
	r.HandleFunc("/{game}/play", s.playHandler)
	r.HandleFunc("/new", s.newGameHandler)
	m.Handle("/", r)

}

func jsonResponse(w http.ResponseWriter, value interface{}) {
	out, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(out))
}
