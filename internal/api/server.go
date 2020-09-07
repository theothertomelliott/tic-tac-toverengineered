package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/turn"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/win"
)

func New(
	turn turn.Controller,
	grid grid.Grid,
	checker win.Checker,
) *Server {
	return &Server{
		turn:    turn,
		grid:    grid,
		checker: checker,
	}
}

type Server struct {
	turn    turn.Controller
	grid    grid.Grid
	checker win.Checker
}

func (s *Server) CreateRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/grid", func(w http.ResponseWriter, req *http.Request) {
		var out [][]*player.Mark
		for i := 0; i < 3; i++ {
			var row []*player.Mark
			for j := 0; j < 3; j++ {
				m, err := s.grid.Mark(grid.Position{X: i, Y: j})
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
	mux.HandleFunc("/player/current", func(w http.ResponseWriter, req *http.Request) {
		current, err := s.turn.NextPlayer()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonResponse(w, current)
	})
	mux.HandleFunc("/winner", func(w http.ResponseWriter, req *http.Request) {
		winner, err := s.checker.Winner()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonResponse(w, winner)
	})
	mux.HandleFunc("/play", func(w http.ResponseWriter, req *http.Request) {
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

		if err := s.turn.TakeTurn(player, pos); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResponse(w, "ok")
	})
}

func jsonResponse(w http.ResponseWriter, value interface{}) {
	out, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(out))
}
