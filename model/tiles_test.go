package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTilesRemove_RemovesCorrectItem(t *testing.T) {
	tiles := NewTilesForNewGame()
	assert.Len(t, tiles, 16)

	// Act
	tiles = tiles.Remove(21)

	// Assert
	assert.Len(t, tiles, 15)
	assert.Equal(t, 22, tiles[0].Value, "Expected that 22 is now the lowest number")

	// Act
	tiles = tiles.Remove(36)

	// Assert
	assert.Len(t, tiles, 14)
	assert.Equal(t, 35, tiles[13].Value, "Expected that 35 is now the highest number")

	// Act -- delete a tile that's not in the collection. Nothing should happen
	tiles = tiles.Remove(21)
	assert.Len(t, tiles, 14)

}
