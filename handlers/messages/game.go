package messages

import (
	"dice-game/model"
	"dice-game/util"
	"time"
)

type NewGameRequest struct {
	Players []string `json:"players"`
}

type GameStateResponse struct {
	Players []PlayerStateResponse `json:"players"`
	Tiles   []model.Tile          `json:"tiles"`
	Stage   string                `json:"stage"`
	Ended   bool                  `json:"ended"`
}

type GameInfoResponse struct {
	Key     string    `json:"key"`
	Players []string  `json:"players"`
	Created time.Time `json:"created"`
}

type PlayerStateResponse struct {
	Name      string      `json:"playerName"`
	TopTile   *model.Tile `json:"topTile,omitempty"`
	GameWorms int         `json:"gameWorms"`
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

func NewGameInfoResponse(game *model.Game) GameInfoResponse {
	return GameInfoResponse{
		Key:     game.Key,
		Players: util.Map(game.Players, func(p *model.Player) string { return p.Name }),
		Created: game.Created,
	}
}

func GetGameInfoResponses(engine *model.Engine) []GameInfoResponse {
	out := make([]GameInfoResponse, 0)
	for _, g := range engine.Games {
		out = append(out, NewGameInfoResponse(g))
	}
	return out
}
