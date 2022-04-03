package main

import (
	"dice-game/handlers"
	"dice-game/model"
	"github.com/gorilla/mux"
	muxlogrus "github.com/pytimer/mux-logrus"
	"net/http"
)

func CreateRouter() *mux.Router {
	// Create handlers
	games := handlers.NewGames(l, model.E)
	moves := handlers.NewMoves(l, model.E)

	// Create router and set routes.
	m := mux.NewRouter()
	m.HandleFunc("/", games.GetInfo).Methods(http.MethodGet)
	m.HandleFunc("/games", games.GetGames).Methods(http.MethodGet)
	m.HandleFunc("/game", games.NewGame).Methods(http.MethodPost)

	// get
	getRouter := m.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/game/{game}/turn", games.GetTurnInfo)
	getRouter.Use(handlers.MiddlewareValidateGame)

	// post
	postRouter := m.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/game/{game}/throw", moves.Throw)
	postRouter.HandleFunc("/game/{game}/pick", moves.Pick)
	postRouter.Use(handlers.MiddlewareValidateGame)

	// Add another logrus instance as middle ware for request logging.
	m.Use(muxlogrus.NewLogger().Middleware)
	return m
}
