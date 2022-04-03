package messages

type TurnStateResponse struct {
	Player   string
	Score    int
	HasWorms bool
	Taken    Dice
}

type ThrowResponse struct {
	Player       string
	Dice         Dice
	AllowedPicks []int
}

type PickDiceRequest struct {
	Pick int
}
