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

	// Validate the game state first.
	if !game.CanThrow() {
		http.Error(w, "throw is only allowed at the start of the turn, or after a pick and only when there are unpicked dice left.", http.StatusMethodNotAllowed)
		return
	}

	// Throw the dice and return the updated turn state.
	m.l.Printf("Throwing dice for game %s, with %d dice left.\n", game.Key, game.CurrentTurnInfo().Remaining)
	throw := game.CurrentTurnInfo().ThrowRemaining()
	m.l.Printf("Thrown: %#v\n", throw)
	GenerateTurnStateResponse(w, game)
}

func (m *Moves) Pick(w http.ResponseWriter, r *http.Request) {
	game := r.Context().Value(GameKey{}).(*model.Game)

	// Validate the game state first.
	if !game.CanPick() {
		http.Error(w, "pick is only allowed after a throw which resulted in unpicked dice", http.StatusMethodNotAllowed)
		return
	}

	// Validate and parse the incoming message
	req, err := ParseRequestBody[messages.PickDiceRequest](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Pick the specified dice, and return the updated turn state.
	m.l.Printf("Game %s | Player %s | Picked %d ", game.Key, game.CurrentTurnInfo().Player.Name, req.Pick)
	game.CurrentTurnInfo().Pick(req.Pick)
	GenerateTurnStateResponse(w, game)
}
