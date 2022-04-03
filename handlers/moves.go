package handlers

import (
	"dice-game/handlers/messages"
	"dice-game/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Moves struct {
	l *logrus.Logger
	e *model.Engine
}

func NewMoves(log *logrus.Logger, engine *model.Engine) *Moves {
	return &Moves{
		l: log,
		e: engine,
	}
}

func (m *Moves) Throw(w http.ResponseWriter, r *http.Request) {
	game := r.Context().Value(GameKey{}).(*model.Game)

	if !game.CanThrow() {
		http.Error(w, "throw is only allowed at the start of the turn, or after a pick and only when there are unpicked dice left.", http.StatusMethodNotAllowed)
		return
	}

	m.l.Printf("Throwing dice for game %s, with %d dice left.\n", game.Key, game.CurrentTurnInfo().Remaining)
	throw := game.CurrentTurnInfo().ThrowRemaining()
	m.l.Printf("Thrown: %#v\n", throw)
	res := &messages.ThrowResponse{
		Player:       game.CurrentTurnInfo().Name,
		Dice:         messages.DiceFromModelDice(throw),
		AllowedPicks: game.CurrentTurnInfo().AllowedPicks(),
	}

	SerializeToResponse(res, w)
}

func (m *Moves) Pick(w http.ResponseWriter, r *http.Request) {
	game := r.Context().Value(GameKey{}).(*model.Game)

	if !game.CanPick() {
		http.Error(w, "pick is only allowed after a throw which resulted in unpicked dice", http.StatusMethodNotAllowed)
		return
	}

	req, err := ParseRequestBody[messages.PickDiceRequest](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	game.CurrentTurnInfo().Pick(req.Pick)

	// Return a 'TurnStateResponse' to show the current status.
	GenerateTurnStateResponse(w, game.CurrentTurnInfo())
}
