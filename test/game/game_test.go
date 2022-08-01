package game_test

import (
	"fmt"
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

func TestOpenCell(t *testing.T) {
	mines := [][]int{
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

	gameOver := func() {
		game, _ := g.NewGame(g.Beginner)
		game.Mines = mines
		x, y := 7, 0
		open := [][]bool{
			{false, false, false, false, false, false, false, true, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
		}
		game.OpenCell(x, y)
		assert.Equal(t, g.Lose, game.State)
		assert.Equal(t, open, game.Open)
	}

	smallOpen := func() {
		game, _ := g.NewGame(g.Beginner)
		game.Mines = mines
		x, y := 0, 4
		open := [][]bool{
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{true, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
		}
		game.OpenCell(x, y)
		assert.Equal(t, g.Playing, game.State)
		assert.Equal(t, open, game.Open)
	}

	bigOpen := func() {
		game, _ := g.NewGame(g.Beginner)
		game.Mines = mines
		x, y := 3, 7
		open := [][]bool{
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, true, true},
			{false, false, false, false, false, false, false, false, true},
			{false, false, false, true, false, false, false, false, true},
			{false, false, false, true, true, false, false, false, true},
			{false, false, false, true, true, true, true, true, true},
			{true, true, true, true, true, true, false, false, false},
			{true, true, true, true, true, true, false, false, false},
		}

		game.OpenCell(x, y)
		assert.Equal(t, 0, game.Mines[y][x])
		assert.Equal(t, g.Playing, game.State)
		assert.Equal(t, open, game.Open)
		fmt.Println(game.Open)
	}

	onFlag := func() {
		game, _ := g.NewGame(g.Beginner)
		open := [][]bool{
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
		}
		flagged := [][]bool{
			{true, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false},
		}
		game.Flagged = flagged
		game.OpenCell(0, 0)
		assert.Equal(t, open, game.Open)
	}

	gameOver()
	smallOpen()
	bigOpen()
	onFlag()
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
