package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTestGame(t *testing.T) *Game {
	e := newEngine()
	game, err := e.NewGame([]string{"dick", "janneke", "lucy"})
	if err != nil {
		t.Fatalf("error creating test-game: %v", err)
	}
	return game
}

func TestGame_CanThrow(t *testing.T) {
	g := getTestGame(t)
	assert.True(t, g.CanThrow())

	d := g.currentTurn.ThrowRemaining()
	assert.False(t, g.CanThrow())

	// just use the first valid value
	g.currentTurn.Pick(d.Values()[0])
	assert.True(t, g.CanThrow())
}

func TestGame_CanPick(t *testing.T) {
	g := getTestGame(t)
	assert.False(t, g.CanPick())

	d := g.currentTurn.ThrowRemaining()
	assert.True(t, g.CanPick())

	// just use the first valid value
	g.currentTurn.Pick(d.Values()[0])
	assert.False(t, g.CanPick())
}

func TestGame_CanTake(t *testing.T) {
	g := getTestGame(t)
	assert.False(t, g.CanTake())

	g.currentTurn.Used[6] = 4
	g.currentTurn.Used[1] = 1
	// we still can't 'take' because the stage is not set to 'picked'
	assert.False(t, g.CanTake())

	// now we should be able to take (21)
	g.currentTurn.Stage = Picked
	assert.True(t, g.CanTake())
}

func TestGame_AllowedTiles(t *testing.T) {
	g := getTestGame(t)
	// Start with an empty set.
	assert.Empty(t, g.AllowedTiles())

	// set the score so the first tile should match
	g.currentTurn.Used[6] = 4 // 4*5 (worm = 5) = 20
	g.currentTurn.Used[1] = 1 // 1

	assert.Equal(t, 21, g.AllowedTiles()[0].Value)

	// higher score still returns 1 allowed tile.
	g.currentTurn.Used[5] = 2 // 2*5 = 10 -> total score is now 31
	assert.Len(t, g.AllowedTiles(), 1, "Still expected 1 element")
	assert.Equal(t, 31, g.AllowedTiles()[0].Value, "Expected the allowed tile matches the score")

	// remove some tiles
	g.Tiles = g.Tiles.Remove(31)
	g.Tiles = g.Tiles.Remove(30)
	assert.Equal(t, 29, g.AllowedTiles()[0].Value, "Expected that the first match is now the highest available tile below 31")

	// now one player has a tile you can steal
	g.Players[1].Tiles.Push(NewTile(31, 3)) // don't select the 'current player'
	assert.Len(t, g.AllowedTiles(), 2, "Expected that both the highest matching tile from the stack and the one we can steal is included.")
	assert.Contains(t, g.AllowedTiles(), NewTile(31, 3))
}

func TestGameThrow_Invalid(t *testing.T) {
	// sets status
	// returns player's last tile to the table

}

func TestGame_TakeFromTable(t *testing.T) {
	// Arrange
	game := getTestGame(t)
	turn := game.CurrentTurnInfo()
	turn.Stage = Picked
	turn.Used = Dice{5: 2, 6: 3} // score = 25

	// Act
	game.Take(25)

	// Assert
	assert.Equal(t, Taken, turn.Stage, "Expected that the stage is set to 'Taken'")
	assert.Len(t, game.Tiles, 15, "Expected that the 'table' tiles are reduced by 1 (16->15)")
	assert.Equal(t, 1, turn.Player.Tiles.Size(), "Expected the current player to now have 1 tile")

}
func TestGame_TakeFromPlayer(t *testing.T) {
	// Arrange
	game := getTestGame(t)
	turn := game.CurrentTurnInfo()
	turn.Stage = Picked
	turn.Used = Dice{5: 2, 6: 3} // score = 25

	// make sure another player has tile 25, and it's no longer on the table.
	game.Tiles.Remove(25)
	game.Players[1].Tiles.Push(NewTile(25, 3))

	// Act
	game.Take(25)

	// Assert
	assert.Equal(t, Taken, turn.Stage, "Expected that the stage is set to 'Taken'")
	assert.Equal(t, 0, game.Players[1].Tiles.Size(), "Expected that the player's tiles are reduced by 1 (1->0)")
	assert.Equal(t, 1, turn.Player.Tiles.Size(), "Expected the current player to now have 1 tile")
}
