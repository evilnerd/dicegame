package model

type Dice map[int]int

func NewDice() Dice {
	return Dice{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0}
}

func (d Dice) Roll(val int) {
	d[val]++
}
