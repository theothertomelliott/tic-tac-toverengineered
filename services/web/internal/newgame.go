package web

import (
	"fmt"
	"net/http"
)

func (s *Server) newGame(w http.ResponseWriter, req *http.Request) {
	matches, err := s.openapiclient.RequestMatchPair(req.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, req, fmt.Sprintf("/%v", matches.O.GameID), http.StatusFound)
}
