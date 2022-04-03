package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTurn_Throw(t *testing.T) {
	p := NewPlayer("dick")
	turn := NewTurn(p)
	d := turn.Throw(5)

	realNum := 0
	for i := 1; i <= 6; i++ {
		realNum += d[i]
	}

	assert.Equal(t, 5, realNum)
}
