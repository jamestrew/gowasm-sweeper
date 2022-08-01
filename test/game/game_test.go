package game_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	g "github.com/jamestrew/gowasm-sweeper/pkg/game"
)

func TestNewGame(t *testing.T) {

	beginnerGame := func() {
		game, error := g.NewGame(g.Beginner)
		assert.Nil(t, error)
		assert.Equal(t, g.Beginner, game.Difficulty)
		assert.Equal(t, 9, game.Width)
		assert.Equal(t, 9, game.Height)
		assert.Equal(t, 10, game.MineCount)
		assert.Equal(t, g.Playing, game.State)
	}

	intermediateGame := func() {
		game, error := g.NewGame(g.Intermediate)
		assert.Nil(t, error)
		assert.Equal(t, g.Intermediate, game.Difficulty)
		assert.Equal(t, 16, game.Width)
		assert.Equal(t, 16, game.Height)
		assert.Equal(t, 40, game.MineCount)
		assert.Equal(t, g.Playing, game.State)
	}

	expertGame := func() {
		game, error := g.NewGame(g.Expert)
		assert.Nil(t, error)
		assert.Equal(t, g.Expert, game.Difficulty)
		assert.Equal(t, 30, game.Width)
		assert.Equal(t, 16, game.Height)
		assert.Equal(t, 99, game.MineCount)
		assert.Equal(t, g.Playing, game.State)
	}

	customGameOk := func() {
		g.CustomWidth = 420
		g.CustomHeight = 69
		g.CustomMineCount = 69
		game, error := g.NewGame(g.Custom)
		assert.Nil(t, error)
		assert.Equal(t, g.Custom, game.Difficulty)
		assert.Equal(t, 420, game.Width)
		assert.Equal(t, 69, game.Height)
		assert.Equal(t, 69, game.MineCount)
		assert.Equal(t, g.Playing, game.State)
	}

	customTooManyMines := func() {
		g.CustomWidth = 3
		g.CustomHeight = 4
		g.CustomMineCount = 69
		game, error := g.NewGame(g.Custom)
		assert.Nil(t, game)
		if assert.Error(t, error) {
			assert.Equal(t, "Mine count (69) too large for given board dimensions (3x4)", error.Error())
		}
	}

	beginnerGame()
	intermediateGame()
	expertGame()
	customGameOk()
	customTooManyMines()
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

	custom := func() {
		g.CustomWidth = 420
		g.CustomHeight = 69
		g.CustomMineCount = 69
		width, height, mines := g.GetBoardParams(g.Custom)
		assert.Equal(t, 420, width)
		assert.Equal(t, 69, height)
		assert.Equal(t, 69, mines)
	}

	beginner()
	intermediate()
	expert()
	custom()
}

func TestGenerateMines(t *testing.T) {
	rand.Seed(69420)
	g.CustomWidth = 4
	g.CustomHeight = 3
	g.CustomMineCount = 4

	game, _ := g.NewGame(g.Custom)
	game.FillMines()

	expected := [][]int{
		{0, 0, 0, 9},
		{9, 9, 0, 0},
		{0, 0, 9, 0},
	}

	assert.Equal(t, expected, game.Mines)
}

func TestCalcAllNeighbors(t *testing.T) {
  customGame := func() {
    g.CustomWidth = 4
    g.CustomHeight = 3
    g.CustomMineCount = 4

    game, _ := g.NewGame(g.Custom)
    game.Mines = [][]int{
      {0, 0, 0, 9},
      {9, 9, 0, 0},
      {0, 0, 9, 0},
    }
    expected := [][]int{
      {2, 2, 2, 9},
      {9, 9, 3, 2},
      {2, 3, 9, 1},
    }
    game.CalcAllNeighbors()
    assert.Equal(t, expected, game.Mines)
  }

  beginnerGame := func() {
    game, _ := g.NewGame(g.Beginner)
    game.Mines = [][]int{
      {0, 0, 0, 0, 0, 0, 0, 9, 0},
      {0, 0, 9, 0, 0, 9, 0, 0, 0},
      {0, 0, 0, 9, 0, 9, 0, 0, 0},
      {0, 0, 0, 0, 0, 9, 0, 0, 0},
      {0, 9, 0, 0, 0, 0, 9, 0, 0},
      {0, 9, 0, 0, 0, 0, 0, 0, 0},
      {0, 0, 0, 0, 0, 0, 0, 0, 0},
      {0, 0, 0, 0, 0, 0, 0, 0, 0},
      {0, 0, 0, 0, 0, 0, 0, 9, 0},
    }
    expected := [][]int{
      {0, 1, 1, 1, 1, 1, 2, 9, 1},
      {0, 1, 9, 2, 3, 9, 3, 1, 1},
      {0, 1, 2, 9, 4, 9, 3, 0, 0},
      {1, 1, 2, 1, 3, 9, 3, 1, 0},
      {2, 9, 2, 0, 1, 2, 9, 1, 0},
      {2, 9, 2, 0, 0, 1, 1, 1, 0},
      {1, 1, 1, 0, 0, 0, 0, 0, 0},
      {0, 0, 0, 0, 0, 0, 1, 1, 1},
      {0, 0, 0, 0, 0, 0, 1, 9, 1},
    }
    game.CalcAllNeighbors()
    assert.Equal(t, expected, game.Mines)
  }

  customGame()
  beginnerGame()
}
