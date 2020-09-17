package api

import "net/http"

func (s *Server) newGameHandler(w http.ResponseWriter, req *http.Request) {
	gameID, err := s.repo.New()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, gameID)
}
