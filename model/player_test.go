package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	p := NewPlayer("dick")
	assert.NotEmpty(t, p)
	assert.Equal(t, "dick", p.Name)
}

func TestPlayer_AddTile(t *testing.T) {
	p := NewPlayer("dick")
	assert.Equal(t, 0, p.Tiles.Size())

	p.Tiles.Push(NewTile(21, 1))

	assert.Equal(t, 1, p.Tiles.Size())
	assert.Equal(t, 21, p.Tiles.Peek().Value)
}

func TestPlayer_Worms(t *testing.T) {
	p := NewPlayer("dick")
	assert.Equal(t, 0, p.Worms())

	p.Tiles.Push(NewTile(21, 1))
	p.Tiles.Push(NewTile(36, 4))

	assert.Equal(t, 5, p.Worms())
}
