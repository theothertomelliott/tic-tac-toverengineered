package web

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

var indexTmpl string

func init() {
	content, err := ioutil.ReadFile("views/index.html")
	if err != nil {
		panic(err)
	}
	indexTmpl = string(content)
}

func (s *Server) index(w http.ResponseWriter, req *http.Request) {
	t, err := template.New("webpage").Parse(indexTmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, struct{}{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
