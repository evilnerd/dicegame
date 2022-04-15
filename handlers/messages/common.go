package messages

import "dice-game/model"

type Dice map[string]int

func DiceFromModelDice(dice model.Dice) Dice {
	if dice.IsEmpty() {
		return nil
	}
	return Dice{"1": dice[1], "2": dice[2], "3": dice[3], "4": dice[4], "5": dice[5], "worms": dice[6]}
}
