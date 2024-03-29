package web

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
)

type Position struct {
	I int32 `json:"i"`
	J int32 `json:"j"`
}

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

	var keyPlayerToken = KeyPlayerTokenO
	if playerParams[0] == "X" {
		keyPlayerToken = KeyPlayerTokenX
	}

	playerTokenCookie, err := req.Cookie(keyPlayerToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	playerTokenBytes, err := base64.StdEncoding.DecodeString(playerTokenCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var pos Position
	err = json.Unmarshal([]byte(posParams[0]), &pos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.apiclient.Play(req.Context(), string(gameID), string(playerTokenBytes), pos.I, pos.J)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, fmt.Sprintf("/%v", gameID), http.StatusFound)
}
