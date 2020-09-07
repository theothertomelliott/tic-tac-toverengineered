package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
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
	r.HandleFunc("/{game}/grid", func(w http.ResponseWriter, req *http.Request) {
		gameID := game.ID(mux.Vars(req)["game"])
		var out [][]*player.Mark
		for i := 0; i < 3; i++ {
			var row []*player.Mark
			for j := 0; j < 3; j++ {
				m, err := s.grid.Mark(gameID, grid.Position{X: i, Y: j})
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				row = append(row, m)
			}
			out = append(out, row)
		}
		jsonResponse(w, out)
	})
	r.HandleFunc("/{game}/player/current", func(w http.ResponseWriter, req *http.Request) {
		gameID := game.ID(mux.Vars(req)["game"])
		current, err := s.turn.NextPlayer(gameID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonResponse(w, current)
	})
	r.HandleFunc("/{game}/winner", func(w http.ResponseWriter, req *http.Request) {
		gameID := game.ID(mux.Vars(req)["game"])
		winner, err := s.checker.Winner(gameID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonResponse(w, winner)
	})
	r.HandleFunc("/{game}/play", func(w http.ResponseWriter, req *http.Request) {
		gameID := game.ID(mux.Vars(req)["game"])
		playerParams, ok := req.URL.Query()["player"]
		if !ok || len(playerParams) == 0 {
			http.Error(w, "player is required", http.StatusInternalServerError)
			return
		}
		player := player.Mark(playerParams[0])

		posParams, ok := req.URL.Query()["pos"]
		if !ok || len(posParams) == 0 {
			http.Error(w, "pos is required", http.StatusInternalServerError)
			return
		}

		var pos grid.Position
		if err := json.Unmarshal([]byte(posParams[0]), &pos); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := s.turn.TakeTurn(gameID, player, pos); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResponse(w, "ok")
	})
	r.HandleFunc("/new", func(w http.ResponseWriter, req *http.Request) {
		gameID, err := s.repo.New()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonResponse(w, gameID)
	})
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
