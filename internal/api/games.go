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

func (s *Server) listGamesHandler(w http.ResponseWriter, req *http.Request) {
	var max int64
	if err := parseParam(req, "max", &max); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	var offset int64
	if err := parseParam(req, "offset", &offset); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	games, err := s.repo.List(max, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonResponse(w, games)
}
