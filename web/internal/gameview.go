package web

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/checker/pkg/win"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
)

func (s *Server) gameview(w http.ResponseWriter, req *http.Request) {
	gameID := game.ID(mux.Vars(req)["game"])
	t, err := template.New("webpage").Parse(gameviewTmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Game       game.ID
		NextPlayer player.Mark
		Winner     win.Result
		Grid       [][]*player.Mark
	}{
		Game: gameID,
	}
	if err := s.client.ApiGet(req.Context(), gameID, "grid", &data.Grid); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := s.client.ApiGet(req.Context(), gameID, "player/current", &data.NextPlayer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := s.client.ApiGet(req.Context(), gameID, "winner", &data.Winner); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

const gameviewTmpl = `
<html>
<head>
	<title>Tic Tac Toe</title>
	<script language="javascript">
		function play(p,x,y){
			{{if not .Winner.Finished }}
			location.href = "/{{.Game}}/play?player=" + p + "&pos={\"X\":" + x + ",\"Y\":" + y + "}";
			{{end}}
		}
	</script>
</head>
<body>
	<h1>Tic Tac Toe</h1>
	<p><a href="/">Home</a></p>
	{{if .Winner.Finished }}
		{{if .Winner.IsDraw}}
			<p>Draw!</p>
		{{else}}
			<p>Winner: {{.Winner.Winner}}</p>
		{{end}}
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
