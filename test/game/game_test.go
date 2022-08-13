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

func TestCountBlankNeighbors(t *testing.T) {
	g.CustomHeight = 3
	g.CustomWidth = 4
	g.CustomMineCount = 1

	noBlankNeighbors := func() {
		game, _ := g.NewGame(g.Custom)
		mines := [][]int{
			{9, 1, 9, 1},
			{1, 1, 1, 1},
			{0, 1, 9, 1},
		}
		game.Mines = mines
		assert.Equal(t, 0, game.CountBlankNeighbors(0, 2))
	}

	hasBlankNeighbors := func() {
		game, _ := g.NewGame(g.Custom)
		mines := [][]int{
			{0, 1, 9, 1},
			{0, 1, 1, 1},
			{0, 1, 9, 1},
		}
		game.Mines = mines
		assert.Equal(t, 1, game.CountBlankNeighbors(0, 0))
		assert.Equal(t, 2, game.CountBlankNeighbors(0, 1))
		assert.Equal(t, 1, game.CountBlankNeighbors(0, 2))
	}

	noBlankNeighbors()
	hasBlankNeighbors()
}

func TestOpenBlankCells(t *testing.T) {
	g.CustomHeight = 3
	g.CustomWidth = 4
	g.CustomMineCount = 1

	noBlankNeighbors := func() {
		game, _ := g.NewGame(g.Custom)
		mines := [][]int{
			{9, 1, 9, 1},
			{1, 1, 1, 1},
			{0, 1, 9, 1},
		}
		open := [][]bool{
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
		}
		game.Mines = mines
		game.OpenBlankCells(0, 2)
		assert.Equal(t, open, game.Open)
	}

	hasBlankNeighbors := func() {
		game, _ := g.NewGame(g.Custom)
		mines := [][]int{
			{0, 1, 9, 1},
			{0, 1, 1, 1},
			{0, 1, 9, 1},
		}
		open := [][]bool{
			{true, false, false, false},
			{true, false, false, false},
			{true, false, false, false},
		}
		game.Mines = mines
		game.OpenBlankCells(0, 2)
		assert.Equal(t, open, game.Open)
	}

	noBlankNeighbors()
	hasBlankNeighbors()
}

func TestFlagCell(t *testing.T) {
	g.CustomHeight = 3
	g.CustomWidth = 4
	g.CustomMineCount = 1

	game, _ := g.NewGame(g.Custom)
	game.Open = [][]bool{
		{true, false, false, false},
		{true, false, false, false},
		{true, false, false, false},
	}
	game.FlagCell(0, 0)
	game.FlagCell(1, 0)
	flagged := [][]bool{
		{false, true, false, false},
		{false, false, false, false},
		{false, false, false, false},
	}
	assert.Equal(t, flagged, game.Flagged)
}

func TestAsArray(t *testing.T) {
	g.CustomWidth = 4
	g.CustomHeight = 3
	g.CustomMineCount = 5
	game, _ := g.NewGame(g.Custom)

	game.Mines = [][]int{
		{9, 1, 9, 1},
		{1, 1, 1, 1},
		{0, 1, 9, 1},
	}

	veryStart := func() {
		expected := [][]int{
			{-1, -1, -1, -1},
			{-1, -1, -1, -1},
			{-1, -1, -1, -1},
		}
		assert.Equal(t, expected, game.AsArray())
	}

	openSome := func() {
		game.Open = [][]bool{
			{false, false, false, false},
			{false, false, false, false},
			{true, false, false, true},
		}
		expected := [][]int{
			{-1, -1, -1, -1},
			{-1, -1, -1, -1},
			{0, -1, -1, 1},
		}
		assert.Equal(t, expected, game.AsArray())
	}

	flagOne := func() {
		game.Flagged = [][]bool{
			{true, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
		}
		expected := [][]int{
			{-2, -1, -1, -1},
			{-1, -1, -1, -1},
			{0, -1, -1, 1},
		}
		assert.Equal(t, expected, game.AsArray())
	}

	openMine := func() {
		game.Open[0][2] = true
		game.State = g.Lose
		expected := [][]int{
			{-2, -1, 9, -1},
			{-1, -1, -1, -1},
			{0, -1, 9, 1},
		}
		assert.Equal(t, expected, game.AsArray())
	}

	veryStart()
	openSome()
	flagOne()
	openMine()
}

func TestAsJson(t *testing.T) {
	g.CustomWidth = 4
	g.CustomHeight = 3
	g.CustomMineCount = 5
	game, _ := g.NewGame(g.Custom)

	game.Mines = [][]int{
		{9, 1, 9, 1},
		{1, 1, 1, 1},
		{0, 1, 9, 1},
	}

	expected := "{\"State\":1,\"Board\":[[-1,-1,-1,-1],[-1,-1,-1,-1],[-1,-1,-1,-1]]}"
	assert.Equal(t, expected, game.AsJson())
}

func TestFoo(t *testing.T) {
	game, _ := g.NewGame(g.Beginner)
	mines := [][]int{
		{0, 1, 9, 1, 0, 0, 0, 0, 0},
		{1, 3, 3, 3, 1, 0, 0, 0, 0},
		{9, 2, 9, 9, 1, 0, 0, 0, 0},
		{1, 2, 2, 2, 1, 0, 0, 0, 0},
		{0, 1, 1, 1, 0, 0, 0, 0, 0},
		{0, 1, 9, 2, 2, 2, 1, 0, 0},
		{0, 1, 1, 2, 9, 9, 2, 0, 0},
		{0, 0, 0, 2, 4, 9, 3, 1, 1},
		{0, 0, 0, 1, 9, 2, 2, 9, 1},
	}
	game.Mines = [][]int{
		{0, 1, 9, 1, 0, 0, 0, 0, 0},
		{1, 3, 3, 3, 1, 0, 0, 0, 0},
		{9, 2, 9, 9, 1, 0, 0, 0, 0},
		{1, 2, 2, 2, 1, 0, 0, 0, 0},
		{0, 1, 1, 1, 0, 0, 0, 0, 0},
		{0, 1, 9, 2, 2, 2, 1, 0, 0},
		{0, 1, 1, 2, 9, 9, 2, 0, 0},
		{0, 0, 0, 2, 4, 9, 3, 1, 1},
		{0, 0, 0, 1, 9, 2, 2, 9, 1},
	}

	game.OpenCell(0, 0)
	assert.Equal(t, mines, game.Mines)
}
