package model

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Game struct {
	Key         string
	Players     []*Player
	Tiles       Tiles
	Created     time.Time
	currentTurn *Turn
	turn        int
}

func (e *Engine) NewGame(players []string) (*Game, error) {
	if len(players) < 2 {
		return nil, fmt.Errorf("at least two players must be specified (currently specified: %d", len(players))
	}

	g := &Game{
		Key:     uuid.New().String(),
		Players: make([]*Player, 0),
		Tiles:   NewTilesForNewGame(),
		Created: time.Now(),
		turn:    0,
	}
	g.createPlayers(players)
	g.resetTurn(g.Players[0])
	e.Games = append(e.Games, g)
	return g, nil
}

// createPlayers creates Players for each supplied name and sets the play order by filling the 'Next' pointer.
func (g *Game) createPlayers(players []string) {
	var prevPlayer *Player
	for _, p := range players {
		player := NewPlayer(p)
		if prevPlayer != nil {
			prevPlayer.Next = player
		}
		g.Players = append(g.Players, player)
		prevPlayer = player
	}
	prevPlayer.Next = g.Players[0]
}

func (g *Game) resetTurn(p *Player) {
	g.turn++
	g.currentTurn = NewTurn(p)
}

func (g *Game) CurrentTurnInfo() *Turn {
	return g.currentTurn
}

func (g *Game) CanPick() bool {
	return g.currentTurn.Stage == Thrown && len(g.currentTurn.AllowedPicks()) > 0
}

func (g *Game) CanThrow() bool {
	return (g.currentTurn.Stage == Start || g.currentTurn.Stage == Picked) &&
		g.currentTurn.Remaining > 0
}

func (g *Game) CanTake() bool {
	return g.currentTurn.Stage == Picked && len(g.AllowedTiles()) > 0
}

func (g *Game) CanNextPlayer() bool {
	return g.currentTurn.Stage == Taken || g.currentTurn.Stage == Invalid
}

func (g *Game) Ended() bool {
	return len(g.Tiles) == 0
}

func (g *Game) AllowedTiles() Tiles {
	ti := Tiles{}

	// check if the player already picked worms
	//   and there are tiles left on the table.
	if !g.currentTurn.HasWorms() || len(g.Tiles) == 0 {
		return ti
	}

	// read from the players and the table.
	for _, p := range g.Players {
		if p.Tiles.Size() > 0 && p.Name != g.currentTurn.Name {
			ti = append(ti, p.Tiles.Peek())
		}
	}
	// read based on what's on the table.
	score := g.currentTurn.Score()
	if score >= g.Tiles[0].Value {
		t := g.Tiles[0]
		for idx := 0; g.Tiles[idx].Value <= score; idx++ {
			t = g.Tiles[idx]
		}
		ti = append(ti, t)
	}

	return ti
}

func (g *Game) NextPlayer() *Turn {
	p := g.currentTurn.Player.Next
	g.currentTurn = NewTurn(p)
	return g.currentTurn
}

func (g *Game) Take(number int) error {
	var tile Tile
	if g.Tiles.Contains(number) {
		// table pick
		tile = g.Tiles.GetByValue(number)
		g.Tiles = g.Tiles.Remove(number)
	} else {
		// check players
		for _, p := range g.Players {
			if p.Name != g.currentTurn.Name && p.Tiles.Peek().Value == number {
				tile = p.Tiles.Pop()
			}
		}
	}
	if tile.Value != 0 {
		g.currentTurn.Player.Tiles.Push(tile)
		g.currentTurn.Stage = Taken
	} else {
		return fmt.Errorf("the specified tile (%d) is not available", number)
	}
	return nil
}
