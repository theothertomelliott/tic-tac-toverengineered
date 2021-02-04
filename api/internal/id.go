package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
)

func (s *Server) verifyID(w http.ResponseWriter, req *http.Request) (game.ID, error) {
	gameID := game.ID(mux.Vars(req)["game"])

	// Record the provided game id for monitoring
	monitoring.AddFieldToSpan(req.Context(), "game_id", gameID)

	// Verify this game exists
	exists, err := s.repo.Exists(req.Context(), gameID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return game.ID(""), err
	}
	if !exists {
		http.Error(w, "game not found", http.StatusNotFound)
		return game.ID(""), fmt.Errorf("game not found: %v", gameID)
	}

	return gameID, nil
}
