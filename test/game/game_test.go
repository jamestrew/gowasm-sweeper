package game_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	g "github.com/jamestrew/gowasm-sweeper/pkg/game"
)

func TestNewGame(t *testing.T) {

	beginnerGame := func() {
		game := g.NewGame(g.Beginner)
		assert.Equal(t, g.Beginner, game.Difficulty)
		assert.Equal(t, 9, game.Width)
		assert.Equal(t, 9, game.Height)
		assert.Equal(t, 10, game.MineCount)
	}

	intermediateGame := func() {
		game := g.NewGame(g.Intermediate)
		assert.Equal(t, g.Intermediate, game.Difficulty)
		assert.Equal(t, 16, game.Width)
		assert.Equal(t, 16, game.Height)
		assert.Equal(t, 40, game.MineCount)
	}

	expertGame := func() {
		game := g.NewGame(g.Expert)
		assert.Equal(t, g.Expert, game.Difficulty)
		assert.Equal(t, 30, game.Width)
		assert.Equal(t, 16, game.Height)
		assert.Equal(t, 99, game.MineCount)
	}

	beginnerGame()
	intermediateGame()
	expertGame()
}

func TestGetBoardParams(t *testing.T) {
	beginner := func() {
		width, height, mines := g.GetBoardParams(g.Beginner)
		assert.Equal(t, 9, width)
		assert.Equal(t, 9, height)
		assert.Equal(t, 10, mines)
	}

	intermediate := func() {
		width, height, mines := g.GetBoardParams(g.Intermediate)
		assert.Equal(t, 16, width)
		assert.Equal(t, 16, height)
		assert.Equal(t, 40, mines)
	}

	expert := func() {
		width, height, mines := g.GetBoardParams(g.Expert)
		assert.Equal(t, 30, width)
		assert.Equal(t, 16, height)
		assert.Equal(t, 99, mines)
	}

	beginner()
	intermediate()
	expert()
}
