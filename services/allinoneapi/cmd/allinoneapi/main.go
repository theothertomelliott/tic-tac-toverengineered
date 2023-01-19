package main

import (
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/env"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/opentelemetry"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/server"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/checker/pkg/win/gridchecker"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/currentturn/pkg/turn/inmemoryturns"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game/inmemoryrepository"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/matchmaker/inmemorymatchmaker"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/matchmaker/unsignedtokens"
	space "github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg/spaceinmemory"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

func main() {
	version.Println()
	cleanup, err := opentelemetry.Setup("allinoneapi")
	if err != nil {
		log.Fatalf("could not configure telemetry: %v", err)
	}
	defer cleanup()

	log.Println("Starting all-in-one api server")

	spaces := [][]space.Space{
		{spaceinmemory.New(), spaceinmemory.New(), spaceinmemory.New()},
		{spaceinmemory.New(), spaceinmemory.New(), spaceinmemory.New()},
		{spaceinmemory.New(), spaceinmemory.New(), spaceinmemory.New()},
	}

	g, err := grid.New(spaces)
	if err != nil {
		log.Fatalf("could not configure grid: %v", err)
	}

	checker := gridchecker.New(g)
	controller := inmemoryturns.New(inmemoryturns.NewCurrentTurn(), g, checker)
	r := inmemoryrepository.New()
	m := inmemorymatchmaker.New(r)

	log.Println("Setting up health check")
	hc, done := startHealthCheck()
	defer close(done)

	apiServer := server.New(
		r,
		m,
		controller,
		g,
		checker,
		&unsignedtokens.UnsignedTokens{},
	)

	e := echo.New()
	e.Pre(healthCheckMiddleWare(hc))
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e.Use(otelecho.Middleware("api"))
	tictactoeapi.RegisterHandlers(e, apiServer)

	port := env.Get("PORT", "8080")
	err = e.Start(fmt.Sprintf(":%v", port))
	if err != nil {
		panic(err)
	}
}

func healthCheckMiddleWare(hc *HealthChecker) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if !hc.Healthy() {
				return echo.ErrServiceUnavailable
			}
			return hf(ctx)
		}
	}
}

func startHealthCheck() (*HealthChecker, chan<- struct{}) {
	hc, err := NewHealthCheck()
	if err != nil {
		panic(err)
	}
	done := make(chan struct{})
	go hc.Run(time.Second, done)
	return hc, done
}
