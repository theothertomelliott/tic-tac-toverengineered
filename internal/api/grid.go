package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
)

func (s *Server) gridHandler(w http.ResponseWriter, req *http.Request) {
	gameID := game.ID(mux.Vars(req)["game"])

	// Verify this game exists
	exists, err := s.repo.Exists(gameID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "game not found", http.StatusNotFound)
		return
	}

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
}
