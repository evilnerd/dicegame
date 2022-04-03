package model

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Game struct {
	Key         string
	Players     []*Player
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
		Created: time.Now(),
		turn:    0,
	}
	g.createPlayers(players)
	g.resetTurn(g.Players[0])
	e.Games = append(e.Games, g)
	return g, nil
}

func (g *Game) createPlayers(players []string) {
	for _, p := range players {
		g.Players = append(g.Players, NewPlayer(p))
	}
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
