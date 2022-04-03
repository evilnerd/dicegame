package model

import "math/rand"

const (
	TOTAL_DICE = 8
)

type Stage int

const (
	Start   Stage = iota // the turn just started.
	Thrown               // player threw the dice and must now pick or declare invalid.
	Picked               // player picked a dice-number to keep. next is either re-throw the remaining dice or take a number from the board/other player.
	Invalid              // the turn is invalid, because of no more throws left and e.g. there are no worms, or the score is too low.
	Stole                // stole a number from another player
	Taken                // taken a number from the board (e.g. '21')
)

type Turn struct {
	*Player
	Stage
	Used      Dice // 1 = 0, 2 = 0, 3 = 1, 4 = 2, 5 = 2, 6 (worm) = 2 -- score == 31
	Remaining int  // the number of dice that are still unpicked. If 0 the player can't throw anymore.
	LastThrow Dice // the last thrown of the player.
}

func NewTurn(p *Player) *Turn {
	return &Turn{
		Player:    p,
		Used:      NewDice(),
		Remaining: TOTAL_DICE,
		Stage:     Start,
	}
}

func (t *Turn) Score() int {

	score := 0
	for i := 1; i <= 5; i++ {
		score += t.Used[i] * i
	}
	// worms
	score += t.Used[6] * 5
	return score
}

func (t *Turn) HasWorms() bool {
	return t.Used[6] > 0
}

func (t *Turn) ThrowRemaining() Dice {
	return t.Throw(t.Remaining)
}

func (t *Turn) Throw(num int) Dice {
	d := NewDice()
	for i := 0; i < num; i++ {
		val := rand.Intn(6) + 1
		d.Roll(val)
	}
	t.Stage = Thrown
	t.LastThrow = d
	return d
}

func (t *Turn) AllowedPicks() []int {
	throw := t.LastThrow
	taken := t.Used
	allowed := make([]int, 0)
	for i := 1; i <= 6; i++ {
		if throw[i] > 0 && taken[i] == 0 {
			allowed = append(allowed, i)
		}
	}
	return allowed
}

func (t *Turn) Pick(number int) {
	throw := t.LastThrow
	amount := throw[number]
	t.Remaining -= amount
	t.Used[number] = amount
	t.Stage = Picked
}
