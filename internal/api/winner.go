package api

import (
	"net/http"
)

func (s *Server) winnerHandler(w http.ResponseWriter, req *http.Request) {
	gameID, err := s.verifyID(w, req)
	if err != nil {
		return
	}

	winner, err := s.checker.Winner(gameID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, winner)
}
