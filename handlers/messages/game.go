package messages

import (
	"dice-game/model"
)

type NewGameRequest struct {
	Players []string
}

type NewGameResponse struct {
	Key string
}
type GameStateResponse struct {
	Players []PlayerStateResponse `json:"players"`
	Tiles   []model.Tile          `json:"tiles"`
	Stage   string                `json:"stage"`
	Ended   bool                  `json:"ended"`
}

type PlayerStateResponse struct {
	Name      string      `json:"player-name"`
	TopTile   *model.Tile `json:"top-tile,omitempty"`
	GameWorms int         `json:"game-worms"`
}

func NewGameStateResponse(game *model.Game) GameStateResponse {

	return GameStateResponse{
		Players: getPlayerStateResponses(game),
		Tiles:   game.Tiles,
		Stage:   game.CurrentTurnInfo().Stage.String(),
		Ended:   game.Ended(),
	}
}

func getPlayerStateResponses(game *model.Game) []PlayerStateResponse {
	out := make([]PlayerStateResponse, 0)
	for _, p := range game.Players {
		out = append(out, newPlayerStateResponse(p))
	}
	return out
}

func newPlayerStateResponse(player *model.Player) PlayerStateResponse {
	res := PlayerStateResponse{
		Name:      player.Name,
		GameWorms: player.Worms(),
	}
	if player.Tiles.Size() > 0 {
		top := player.Tiles.Peek()
		res.TopTile = &top
	}
	return res
}
