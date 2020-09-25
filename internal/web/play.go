package web

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
)

func (s *Server) play(w http.ResponseWriter, req *http.Request) {
	gameID := game.ID(mux.Vars(req)["game"])
	playerParams, ok := req.URL.Query()["player"]
	if !ok || len(playerParams) == 0 {
		http.Error(w, "player is required", http.StatusInternalServerError)
		return
	}
	posParams, ok := req.URL.Query()["pos"]
	if !ok || len(posParams) == 0 {
		http.Error(w, "pos is required", http.StatusInternalServerError)
		return
	}

	resp, err := s.client.Get(req.Context(), gameID, fmt.Sprintf("play?player=%v&pos=%v", playerParams[0], posParams[0]))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		var msg string = string(body)
		if err != nil {
			msg = err.Error()
		}
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, req, fmt.Sprintf("/%v", gameID), http.StatusFound)
}
