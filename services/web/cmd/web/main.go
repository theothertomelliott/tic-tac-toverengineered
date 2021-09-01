package main

import (
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
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/apiclient"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi/tictactoeapiclient"
	web "github.com/theothertomelliott/tic-tac-toverengineered/services/web/internal"
)

func getAPIBaseURL() string {
	if apiBaseURL := os.Getenv("API_BASE_URL"); apiBaseURL != "" {
		return apiBaseURL
	}
	return "http://localhost:8081"
}

func getOpenAPIBaseURL() string {
	if apiBaseURL := os.Getenv("OPENAPI_BASE_URL"); apiBaseURL != "" {
		return apiBaseURL
	}
	return "http://localhost:8094"
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

	openapiClient, err := tictactoeapi.NewClientWithResponses(getOpenAPIBaseURL(), tictactoeapi.WithHTTPClient(client))
	if err != nil {
		log.Fatal(err)
	}

	server := web.New(
		apiclient.New(
			getAPIBaseURL(),
			client,
		),
		tictactoeapiclient.New(openapiClient),
	)

	port := env.Get("PORT", "8080")
	log.Printf("Listening on port :%v\n", port)

	m := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public"))
	m.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	prefix := os.Getenv("ROUTE_PREFIX")
	if prefix != "" && prefix != "/" {
		m = m.PathPrefix(prefix).Subrouter()
	}
	server.AddRoutes(m)

	http.ListenAndServe(fmt.Sprintf(":%v", port), m)
}
