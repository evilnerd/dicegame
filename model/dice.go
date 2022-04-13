package model

type Dice map[int]int

func NewDice() Dice {
	return Dice{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0}
}

func (d Dice) Roll(val int) {
	d[val]++
}

func (d Dice) IsEmpty() bool {
	for _, v := range d {
		if v > 0 {
			return false
		}
	}
	return true
}

// Values returns a list of 'eyes' that were rolled.
func (d Dice) Values() []int {
	out := make([]int, 0)
	for eyes, occ := range d {
		if occ > 0 {
			out = append(out, eyes)
		}
	}
	return out
}
