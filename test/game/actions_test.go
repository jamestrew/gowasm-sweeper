package game_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	g "github.com/jamestrew/gowasm-sweeper/pkg/game"
)

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
			{false, false, false, false, false, false, true, true, true},
			{false, false, false, false, false, false, true, true, true},
			{false, false, true, true, true, false, true, true, true},
			{false, false, true, true, true, true, false, true, true},
			{false, false, true, true, true, true, true, true, true},
			{true, true, true, true, true, true, true, true, true},
			{true, true, true, true, true, true, true, true, true},
			{true, true, true, true, true, true, true, false, false},
		}

		game.OpenCell(x, y)
		assert.Equal(t, 0, game.Mines[y][x])
		assert.Equal(t, g.Playing, game.State)
		assert.Equal(t, open, game.Open)
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

	winning := func() {
		g.CustomWidth = 4
		g.CustomHeight = 3
		g.CustomMineCount = 5
		game, _ := g.NewGame(g.Custom)

		game.Mines = [][]int{
			{9, 1, 9, 1},
			{1, 1, 1, 1},
			{0, 1, 9, 1},
		}
		game.Open = [][]bool{
			{false, true, false, true},
			{true, true, true, true},
			{false, true, false, true},
		}
		expected := [][]bool{
			{false, true, false, true},
			{true, true, true, true},
			{true, true, false, true},
		}

		flagged := [][]bool{
			{true, false, true, false},
			{false, false, false, false},
			{false, false, true, false},
		}
		game.OpenCell(0, 2)
		assert.Equal(t, expected, game.Open)
		assert.Equal(t, g.Win, game.State)
		assert.Equal(t, flagged, game.Flagged)
	}

	notPlaying := func() {
		g.CustomWidth = 4
		g.CustomHeight = 3
		g.CustomMineCount = 5
		game, _ := g.NewGame(g.Custom)

		game.Mines = [][]int{
			{9, 1, 9, 1},
			{1, 1, 1, 1},
			{0, 1, 9, 1},
		}
		game.Open = [][]bool{
			{false, true, false, true},
			{true, true, true, true},
			{false, true, false, true},
		}
		expected := [][]bool{
			{false, true, false, true},
			{true, true, true, true},
			{false, true, false, true},
		}
		game.State = g.Win
		game.OpenCell(0, 2)
		assert.Equal(t, expected, game.Open)

		game.State = g.Lose
		game.OpenCell(0, 2)
		assert.Equal(t, expected, game.Open)
	}

	gameOver()
	smallOpen()
	bigOpen()
	onFlag()
	winning()
	notPlaying()
}
