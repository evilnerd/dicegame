package messages

import "dice-game/model"

type NewGameRequest struct {
	Players []string
}

type NewGameResponse struct {
	Key string
}

type Dice map[string]int

func DiceFromModelDice(dice model.Dice) Dice {
	return Dice{"1": dice[1], "2": dice[2], "3": dice[3], "4": dice[4], "5": dice[5], "worms": dice[6]}
}
