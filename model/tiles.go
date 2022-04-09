package model

type Tile struct {
	Value int
	Worms int
}

type Tiles []Tile

type PlayerTiles = Stack[Tile]

func NewTile(val int, worms int) Tile {
	return Tile{
		Value: val,
		Worms: worms,
	}
}

func NewTiles() Tiles {
	return Tiles{}
}

func NewTilesForNewGame() Tiles {
	return Tiles{
		NewTile(21, 1),
		NewTile(22, 1),
		NewTile(23, 1),
		NewTile(24, 1),
		NewTile(25, 2),
		NewTile(26, 2),
		NewTile(27, 2),
		NewTile(28, 2),
		NewTile(29, 3),
		NewTile(30, 3),
		NewTile(31, 3),
		NewTile(32, 3),
		NewTile(33, 4),
		NewTile(34, 4),
		NewTile(35, 4),
		NewTile(36, 4),
	}
}

func NewPlayerTiles() *PlayerTiles {
	return &PlayerTiles{}
}

func (t Tiles) AsIntSlice() []int {
	out := make([]int, len(t))
	for i, ti := range t {
		out[i] = ti.Value
	}
	return out

}

func (t Tiles) Remove(val int) Tiles {
	for i := 0; i < len(t); i++ {
		if t[i].Value == val {
			return append(t[:i], t[i+1:]...)
		}
	}
	return t
}
