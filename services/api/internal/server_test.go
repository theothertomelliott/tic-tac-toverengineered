package api_test

import (
	"testing"

	api "github.com/theothertomelliott/tic-tac-toverengineered/services/api/internal"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/checker/pkg/win/gridchecker"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/currentturn/pkg/turn/inmemoryturns"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game/inmemoryrepository"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/grid/pkg/grid"
	space "github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/space/pkg/spaceinmemory"
)

func TestServer(t *testing.T) {
	grid, err := grid.New([][]space.Space{
		{spaceinmemory.New(), spaceinmemory.New(), spaceinmemory.New()},
		{spaceinmemory.New(), spaceinmemory.New(), spaceinmemory.New()},
		{spaceinmemory.New(), spaceinmemory.New(), spaceinmemory.New()},
	})
	if err != nil {
		t.Fatal(err)
	}
	checker := gridchecker.New(grid)

	server := api.New(
		inmemoryrepository.New(),
		inmemoryturns.New(
			inmemoryturns.NewCurrentTurn(),
			grid,
			checker,
		),
		grid,
		checker,
	)
	t.Log(server)

}
