package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi/tictactoeapiclient"
)

func HealthCheckMiddleware(ctx context.Context, client *tictactoeapiclient.Client) mux.MiddlewareFunc {
	var apiHealthy bool

	go func() {
		for {
			_, err := client.Index(ctx, nil, nil)
			if err == nil {
				apiHealthy = true
			}
			time.Sleep(time.Second)
		}
	}()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !apiHealthy {
				serveWaitingForApiMessage(w, r)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

var waitingForApiTmpl string

func init() {
	content, err := ioutil.ReadFile("views/waitingforapi.html")
	if err != nil {
		panic(err)
	}
	waitingForApiTmpl = string(content)
}

func serveWaitingForApiMessage(w http.ResponseWriter, req *http.Request) {
	t, err := template.New("webpage").Parse(waitingForApiTmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusServiceUnavailable)
	err = t.Execute(w, struct{}{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
