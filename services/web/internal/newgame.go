package web

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func (s *Server) newGame(w http.ResponseWriter, req *http.Request) {
	matches, err := s.openapiclient.RequestMatchPair(req.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  KeyPlayerTokenX,
		Value: base64.StdEncoding.EncodeToString([]byte(matches.X.Token)),
	})
	http.SetCookie(w, &http.Cookie{
		Name:  KeyPlayerTokenO,
		Value: base64.StdEncoding.EncodeToString([]byte(matches.O.Token)),
	})
	http.Redirect(w, req, fmt.Sprintf("/%v", matches.O.GameID), http.StatusFound)
}
