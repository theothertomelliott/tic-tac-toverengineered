package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
)

func (s *Server) winnerHandler(w http.ResponseWriter, req *http.Request) {
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

	winner, err := s.checker.Winner(gameID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, winner)
}
