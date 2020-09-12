package web

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	t, err := template.New("webpage").Parse(indexTmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct{}{}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

const indexTmpl = `
<html>
<head>
<title>Tic Tac Toe</title>
</head>
<body>
	<h1>Tic Tac Toe</h1>
	<p>Welcome</p>
</body>
</html>
`
