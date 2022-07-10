package messages

import (
	"dice-game/model"
)

type TurnStateResponse struct {
	Player        string `json:"playerName"`
	Score         int    `json:"score"`
	HasWorms      bool   `json:"hasWorms"`
	Taken         Dice   `json:"taken,omitempty"`
	Stage         string `json:"stage"`
	ThrowResponse `json:"throwResponse,omitempty"`
}

type ThrowResponse struct {
	Dice         Dice  `json:"dice,omitempty"`
	AllowedPicks []int `json:"allowedPicks,omitempty"`
	AllowedTiles []int `json:"allowedTiles,omitempty"`
}

type PickDiceRequest struct {
	Pick int `json:"pick"`
}

type TakeTileRequest struct {
	Tile int `json:"tile"`
}

func NewTurnStateResponse(game *model.Game) TurnStateResponse {
	turnInfo := game.CurrentTurnInfo()

	stateRes := TurnStateResponse{
		Player:   turnInfo.Name,
		Score:    turnInfo.Score(),
		HasWorms: turnInfo.HasWorms(),
		Stage:    turnInfo.Stage.String(),
		Taken:    DiceFromModelDice(turnInfo.Used),
	}

	if turnInfo.Stage == model.Thrown || turnInfo.Stage == model.Picked {
		throwRes := ThrowResponse{
			Dice:         DiceFromModelDice(turnInfo.LastThrow),
			AllowedPicks: turnInfo.AllowedPicks(),
			AllowedTiles: game.AllowedTiles().AsIntSlice(),
		}
		stateRes.ThrowResponse = throwRes
	}

	return stateRes
}
