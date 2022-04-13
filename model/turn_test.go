package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTestTurn() *Turn {
	p := NewPlayer("dick")
	turn := NewTurn(p)
	return turn
}

func TestTurn_Throw(t *testing.T) {
	turn := getTestTurn()
	d := turn.Throw(5)

	realNum := 0
	for i := 1; i <= 6; i++ {
		realNum += d[i]
	}

	assert.Equal(t, 5, realNum)
}

func TestTurnThrow_SetsGameStage(t *testing.T) {
	turn := getTestTurn()
	_ = turn.ThrowRemaining()

	assert.Equal(t, Thrown, turn.Stage, "Expected the stage to be set to 'Thrown'")
}
