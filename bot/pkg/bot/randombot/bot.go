package randombot

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/theothertomelliott/tic-tac-toverengineered/bot/pkg/bot"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid"
)

func New() bot.Bot {
	s := rand.NewSource(time.Now().Unix())
	return &randombot{
		random: rand.New(s),
	}
}

type randombot struct {
	random *rand.Rand
}

func (r *randombot) Move(mark player.Mark, state [][]*player.Mark) (grid.Position, error) {
	var valid []grid.Position
	for i, row := range state {
		for j, m := range row {
			if m == nil {
				valid = append(
					valid,
					grid.Position{
						X: i,
						Y: j,
					},
				)
			}
		}
	}
	if len(valid) == 0 {
		return grid.Position{}, fmt.Errorf("no valid moves")
	}
	return valid[r.random.Intn(len(valid))], nil
}
