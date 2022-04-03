package testify

import (
	"dice-game/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BasicGameSuite struct {
	suite.Suite
	game    *model.Game
	players []string
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &BasicGameSuite{})
}

func (s *BasicGameSuite) SetupSuite() {
	s.players = []string{"Dick", "Janneke", "Lucy"}
	g, err := model.E.NewGame(s.players)
	if err != nil {
		s.FailNow("error creating new game: %v", err)
	}
	s.game = g
}

func (s *BasicGameSuite) TestCreateGame() {
	s.T().Run("Create Game", func(t *testing.T) {
		assert.NotEmpty(t, s.game.Created)
		assert.NotEmpty(t, s.game.Key)
		assert.NotNil(t, s.game.CurrentTurnInfo())
	})
	s.T().Run("Inspect First Turn Info", func(t *testing.T) {
		turn := s.game.CurrentTurnInfo()
		assert.NotEmpty(t, turn.Name)
		assert.False(t, turn.HasWorms())
		for _, v := range turn.Used {
			assert.Equal(t, 0, v)
		}
	})
}

func (s *BasicGameSuite) TestThrow() {
}
