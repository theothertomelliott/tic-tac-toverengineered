package api

import (
	"net/http"
)

func (s *Server) gridHandler(w http.ResponseWriter, req *http.Request) {
	gameID, err := s.verifyID(w, req)
	if err != nil {
		return
	}
	out, err := s.grid.State(req.Context(), gameID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, out)
}
