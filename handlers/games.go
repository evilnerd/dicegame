package handlers

import (
	"dice-game/handlers/messages"
	"dice-game/model"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Games struct {
	l *logrus.Logger
	e *model.Engine
}

type GameKey struct{}

func NewGames(logger *logrus.Logger, engine *model.Engine) *Games {
	return &Games{
		l: logger,
		e: engine,
	}
}

func (g *Games) GetInfo(w http.ResponseWriter, r *http.Request) {
	g.l.Println("Returning system information")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("System running OK"))
}

func (g *Games) GetGames(w http.ResponseWriter, r *http.Request) {
	g.l.Println("Listing Games")
	SerializeToResponse(messages.GetGameInfoResponses(g.e), w)
}

func (g *Games) NewGame(w http.ResponseWriter, r *http.Request) {
	req, err := ParseRequestBody[messages.NewGameRequest](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	game, err := g.e.NewGame(req.Players)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := messages.NewGameInfoResponse(game)
	SerializeToResponse(res, w)
}

func (g *Games) NextPlayer(w http.ResponseWriter, r *http.Request) {
	game := r.Context().Value(GameKey{}).(*model.Game)
	if !game.CanNextPlayer() {
		http.Error(w, "the current turn has not yet ended in either taking a tile or an invalid throw", http.StatusMethodNotAllowed)
		return
	}
	// start a new turn for the next player
	game.NextPlayer()
	GenerateTurnStateResponse(w, game)
}

func (g *Games) GetTurnInfo(w http.ResponseWriter, r *http.Request) {
	game := r.Context().Value(GameKey{}).(*model.Game)
	GenerateTurnStateResponse(w, game)
}

func (g *Games) GetGameInfo(w http.ResponseWriter, r *http.Request) {
	game := r.Context().Value(GameKey{}).(*model.Game)
	GenerateGameStateResponse(w, game)
}

func GenerateTurnStateResponse(w http.ResponseWriter, game *model.Game) {
	res := messages.NewTurnStateResponse(game)
	err := SerializeToResponse(res, w)
	if err != nil {
		http.Error(w, fmt.Sprintf("error encoding result: %v", err), http.StatusInternalServerError)
	}
}

func GenerateGameStateResponse(w http.ResponseWriter, game *model.Game) {
	res := messages.NewGameStateResponse(game)
	err := SerializeToResponse(res, w)
	if err != nil {
		http.Error(w, fmt.Sprintf("error encoding result: %v", err), http.StatusInternalServerError)
	}
}
