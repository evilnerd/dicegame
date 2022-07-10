package handlers

import (
	"dice-game/handlers/messages"
	"dice-game/model"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGames_GetGames(t *testing.T) {

	// Arrange
	e := model.NewEngine()
	g := NewGames(logrus.New(), e)

	req, err := http.NewRequest("GET", "/games", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(g.GetGames)

	// Act
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code, "Expected to get an OK status.")
	assert.JSONEq(t, "[]", rr.Body.String(), "The handler returned an unexpected body.")

	// Act - Test 2
	game, _ := e.NewGame([]string{"lucy", "janneke"})
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Assert
	expected := "[ " + AsJSON(messages.NewGameInfoResponse(game)) + " ]"
	assert.Equal(t, http.StatusOK, rr.Code, "Expected to get an OK status.")
	assert.JSONEq(t, expected, rr.Body.String(), "The handler returned an unexpected body.")

}

func TestGames_NewGame(t *testing.T) {
	// Arrange
	e := model.NewEngine()
	g := NewGames(logrus.New(), e)

	req, err := http.NewRequest("POST", "/game", strings.NewReader(`{ "players": ["dick", "janneke", "lucy"] }`))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(g.NewGame)

	// Act
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	res := rr.Body.String()

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code, "Expected to get an OK status.")

	// Try to unmarshal to a GameInfo type
	gameInfo := &messages.GameInfoResponse{}
	err = json.Unmarshal([]byte(res), &gameInfo)
	assert.NoError(t, err, "Expected that the resulting message could be unmarshalled into a GameInfoResponse struct")
	assert.Lenf(t, gameInfo.Players, 3, "Expected the GameInfoResponse to contain the information about the 3 players we used to create the game")
	assert.NotEmpty(t, gameInfo.Key, "Expected the GameInfoResponse to contain a key for the new game.")
}
