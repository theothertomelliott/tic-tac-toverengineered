package api

import (
	"net/http"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/param"
)

func (s *Server) newGameHandler(w http.ResponseWriter, req *http.Request) {
	gameID, err := s.repo.New(req.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, gameID)
}

func (s *Server) listGamesHandler(w http.ResponseWriter, req *http.Request) {
	var max int64
	if err := param.Parse(req, "max", &max, param.ParseOptions{Default: 10}); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	var offset int64
	if err := param.Parse(req, "offset", &offset, param.ParseOptions{Default: 0}); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	games, err := s.repo.List(req.Context(), max, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonResponse(w, games)
}
