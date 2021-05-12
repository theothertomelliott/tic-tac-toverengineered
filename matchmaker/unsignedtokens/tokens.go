package unsignedtokens

import (
	"encoding/json"
	"fmt"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/matchmaker"
)

var _ matchmaker.TokenCreator = &UnsignedTokens{}
var _ matchmaker.TokenValidator = &UnsignedTokens{}

type UnsignedTokens struct {
}

type tokenBody struct {
	Game   game.ID
	Player player.Mark
}

func (u *UnsignedTokens) Create(game game.ID, player player.Mark) (string, error) {
	var t = tokenBody{
		Game:   game,
		Player: player,
	}
	j, err := json.Marshal(t)
	if err != nil {
		return "", fmt.Errorf("creating token: %w", err)
	}
	return string(j), nil
}

func (u *UnsignedTokens) Validate(token string) (game.ID, player.Mark, error) {
	var t = tokenBody{}
	err := json.Unmarshal([]byte(token), &t)
	if err != nil {
		return "", "", fmt.Errorf("creating token: %w", err)
	}
	return t.Game, t.Player, nil
}
