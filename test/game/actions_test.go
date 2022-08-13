package game_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	g "github.com/jamestrew/gowasm-sweeper/pkg/game"
)

func foo(actual [][]bool) [][]int {
	ret := [][]int{
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

	for i, row := range actual {
		for j, cell := range row {
			if cell {
				ret[j][i] = 1
			} else {
				ret[j][i] = 0
			}
		}
	}
	return ret
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
		fmt.Println(foo(game.Open))
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
		fmt.Println(foo(game.Open))
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
