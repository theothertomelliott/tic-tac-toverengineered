package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/turn"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/win"
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

func parseParam(req *http.Request, key string, value interface{}) error {
	params, ok := req.URL.Query()[key]
	if !ok || len(params) == 0 {
		return fmt.Errorf("%q is required", key)
	}
	err := json.Unmarshal([]byte(params[0]), value)
	if err != nil {
		// Allow for string-type inputs that aren't quote delimited
		if err2 := json.Unmarshal([]byte(fmt.Sprintf("\"%v\"", params[0])), value); err2 != nil {
			return fmt.Errorf("parsing %q: %w", key, err)
		}
	}
	return nil
}
