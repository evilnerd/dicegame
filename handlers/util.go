package handlers

import (
	"context"
	"dice-game/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func ParseRequestBody[T any](r *http.Request) (T, error) {
	output := new(T)
	d := json.NewDecoder(r.Body)
	err := d.Decode(&output)
	if err != nil {
		return *output, err
	}
	return *output, nil
}

func SerializeToResponse(in any, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w)
	return e.Encode(in)
}

func AsJSON(in any) string {
	b, _ := json.Marshal(in)
	return string(b)
}

func MiddlewareValidateGame(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key, ok := vars["game"]

		if !ok {
			http.Error(w, "no game key specified", http.StatusBadRequest)
			return
		}

		game, err := model.E.Game(key)

		if err != nil {
			http.Error(w, fmt.Sprintf("game with key %s does not exist", key), http.StatusBadRequest)
			return
		}

		// add the game to the context
		ctx := context.WithValue(r.Context(), GameKey{}, game)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
