package model

import "fmt"

var (
	E = newEngine()
)

type Engine struct {
	Games []*Game
}

// newEngine returns a new Engine. This is called only to initialize the global var E.
func newEngine() *Engine {
	return &Engine{
		Games: make([]*Game, 0),
	}
}

// Game returns the game with the specified key.
func (e *Engine) Game(key string) (*Game, error) {
	for _, g := range e.Games {
		if g.Key == key {
			return g, nil
		}
	}
	return nil, fmt.Errorf("game with key %s not found", key)
}
