package main

import (
	"github.com/labstack/echo/v4"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/server"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game/inmemoryrepository"
)

func main() {
	apiServer := server.New(inmemoryrepository.New(), nil, nil, nil)
	e := echo.New()
	tictactoeapi.RegisterHandlers(e, apiServer)
	err := e.Start(":8080")
	if err != nil {
		panic(err)
	}
}
