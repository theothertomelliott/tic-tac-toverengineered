package randombot

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/bot/pkg/bot"
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

func (r *randombot) Move(mark player.Mark, state [][]string) (tictactoeapi.Position, error) {
	var valid []tictactoeapi.Position
	for i, row := range state {
		for j, m := range row {
			if m == "" {
				valid = append(
					valid,
					tictactoeapi.Position{
						I: int32(i),
						J: int32(j),
					},
				)
			}
		}
	}
	if len(valid) == 0 {
		return tictactoeapi.Position{}, fmt.Errorf("no valid moves")
	}
	return valid[r.random.Intn(len(valid))], nil
}
