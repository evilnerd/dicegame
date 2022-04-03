package handlers

import (
	"dice-game/handlers/messages"
	"dice-game/model"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Games struct {
	l *logrus.Logger
	e *model.Engine
}

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
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	err := e.Encode(g.e.Games)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		g.l.Printf("Error encoding games: %v\n", err)
	}
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
	res := &messages.NewGameResponse{
		Key: game.Key,
	}
	SerializeToResponse(res, w)
}

func (g *Games) GetTurnInfo(w http.ResponseWriter, r *http.Request) {
	game := r.Context().Value(GameKey{}).(*model.Game)
	turn := game.CurrentTurnInfo()

	GenerateTurnStateResponse(w, turn)
}

func GenerateTurnStateResponse(w http.ResponseWriter, turn *model.Turn) {
	res := &messages.TurnStateResponse{
		Player:   turn.Name,
		Score:    turn.Score(),
		HasWorms: turn.HasWorms(),
		Taken:    messages.DiceFromModelDice(turn.Used),
	}
	err := SerializeToResponse(res, w)
	if err != nil {
		http.Error(w, fmt.Sprintf("error encoding result: %v", err), http.StatusInternalServerError)
	}
}

type GameKey struct{}
