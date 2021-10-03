package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/env"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/defaultmonitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi/tictactoeapiclient"
	web "github.com/theothertomelliott/tic-tac-toverengineered/services/web/internal"
)

func getAPIBaseURL() string {
	if apiBaseURL := os.Getenv("API_BASE_URL"); apiBaseURL != "" {
		return apiBaseURL
	}
	return "http://localhost:8080"
}

func main() {
	version.Println()
	defaultmonitoring.Init("web")
	defer monitoring.Close()

	log.Println("Starting web")
	client := &http.Client{
		Transport: monitoring.WrapHTTPTransport(http.DefaultTransport),
		Timeout:   time.Second * 5,
	}

	apiClient, err := tictactoeapi.NewClientWithResponses(getAPIBaseURL(), tictactoeapi.WithHTTPClient(client))
	if err != nil {
		log.Fatal(err)
	}

	tttClient := tictactoeapiclient.New(apiClient)
	server := web.New(
		tttClient,
	)

	port := env.Get("PORT", "8080")
	log.Printf("Listening on port :%v\n", port)

	m := mux.NewRouter()

	m.Use(HealthCheckMiddleware(context.Background(), tttClient))

	fs := http.FileServer(http.Dir("./public"))
	m.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	prefix := os.Getenv("ROUTE_PREFIX")
	if prefix != "" && prefix != "/" {
		m = m.PathPrefix(prefix).Subrouter()
	}
	server.AddRoutes(m)

	http.ListenAndServe(fmt.Sprintf(":%v", port), m)
}
