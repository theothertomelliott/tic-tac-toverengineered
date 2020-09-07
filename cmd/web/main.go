package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{game}", func(w http.ResponseWriter, req *http.Request) {
		gameID := game.ID(mux.Vars(req)["game"])
		t, err := template.New("webpage").Parse(tmpl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			Game       game.ID
			NextPlayer player.Mark
			Winner     *player.Mark
			Grid       [][]*player.Mark
		}{
			Game: gameID,
		}
		if err := apiGet("http://localhost:8081/%v/grid", gameID, &data.Grid); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := apiGet("http://localhost:8081/%v/player/current", gameID, &data.NextPlayer); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := apiGet("http://localhost:8081/%v/winner", gameID, &data.Winner); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	r.HandleFunc("/{game}/play", func(w http.ResponseWriter, req *http.Request) {
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

		resp, err := http.Get(fmt.Sprintf("http://localhost:8081/%v/play?player=%v&pos=%v", gameID, playerParams[0], posParams[0]))
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
	})

	http.ListenAndServe(":8080", r)
}

func apiGet(urlFmt string, g game.ID, out interface{}) error {
	resp, err := http.Get(fmt.Sprintf(urlFmt, g))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, out)
	if err != nil {
		return err
	}
	return nil
}

const tmpl = `
<html>
<head>
	<title>Tic Tac Toe</title>
	<script language="javascript">
		function play(p,x,y){
			{{if not .Winner}}
			location.href = "/{{.Game}}/play?player=" + p + "&pos={\"X\":" + x + ",\"Y\":" + y + "}";
			{{end}}
		}
	</script>
</head>
<body>
	<h1>Tic Tac Toe</h1>
	{{if .Winner}}
		<p>Winner: {{.Winner}}</p>
	{{else}}
		<p>Next Player: {{.NextPlayer}}</p>
	{{end}}

	{{range $i, $r := .Grid}}
		<p>
			{{range $j, $s := $r}}
				<button type="button" onclick="javascript: play('{{$.NextPlayer}}',{{$i}},{{$j}});">
					{{if not $s }}
						&nbsp;&nbsp;
					{{else}}
						{{$s}}
					{{end}}
				</button>&nbsp;&nbsp;
			{{end}}
		</p>
	{{else}}
		<p>No grid</p>
	{{end}}
</body>
</html>
`
