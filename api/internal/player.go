package api

import (
	"net/http"
)

func (s *Server) currentPlayerHandler(w http.ResponseWriter, req *http.Request) {
	gameID, err := s.verifyID(w, req)
	if err != nil {
		return
	}

	current, err := s.turn.NextPlayer(req.Context(), gameID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, current)
}
