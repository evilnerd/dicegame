package model

type Player struct {
	Name  string
	Tiles *Stack[Tile]
	Next  *Player
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:  name,
		Tiles: NewPlayerTiles(),
	}
}

func (p *Player) AddTile(t Tile) {
	p.Tiles.Push(t)
}

func (p Player) Worms() int {
	w := 0
	for _, t := range p.Tiles.entries {
		w += t.Worms
	}
	return w
}
