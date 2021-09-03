package web

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
)

var gameviewTmpl string

func init() {
	content, err := ioutil.ReadFile("views/game.html")
	if err != nil {
		panic(err)
	}
	gameviewTmpl = string(content)
}

func (s *Server) gameview(w http.ResponseWriter, req *http.Request) {
	gameID := game.ID(mux.Vars(req)["game"])
	t, err := template.New("webpage").Parse(gameviewTmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Game       game.ID
		NextPlayer string
		Finished   bool
		Draw       bool
		Winner     string
		Grid       [][]string
	}{
		Game: gameID,
	}

	data.Grid, err = s.openapiclient.GameGrid(req.Context(), string(gameID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.NextPlayer, err = s.openapiclient.CurrentPlayer(req.Context(), string(gameID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	winner, err := s.openapiclient.Winner(req.Context(), string(gameID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data.Draw = winner.Draw != nil && *winner.Draw
	data.Finished = data.Draw || winner.Winner != nil
	if winner.Winner != nil {
		data.Winner = *winner.Winner
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
