package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	game := NewGame(Beginner)

	assert.Equal(t, game.difficulty.width, 9)
	assert.Equal(t, game.difficulty.height, 9)
	assert.Equal(t, game.state, Playing)
	assert.Equal(t, game.state, Playing)

	mineCount := 0
	for _, row := range *(game.board) {
		for _, cell := range row {
			if cell == Mine {
				mineCount++
			}
		}
	}
	assert.Equal(t, 9, mineCount)
}

func TestCalcNeighbors(t *testing.T) {

	b := Board{
		{Closed, Closed, Closed},
		{Closed, Mine, Closed},
		{Closed, Closed, Closed},
	}

	for i, row := range b {
		for j := range row {
			if i == j {
				continue
			}
			assert.Equal(t, b.CalcNeighbors(i, j), 1)
		}
	}

	b = Board{
		{Mine, Closed, Closed},
		{Closed, Closed, Closed},
		{Closed, Closed, Closed},
	}
	assert.Equal(t, 0, b.CalcNeighbors(0, 0))
	assert.Equal(t, 1, b.CalcNeighbors(0, 1))
	assert.Equal(t, 1, b.CalcNeighbors(1, 0))
	assert.Equal(t, 1, b.CalcNeighbors(1, 1))
	assert.Equal(t, 0, b.CalcNeighbors(0, 2))
	assert.Equal(t, 0, b.CalcNeighbors(2, 0))

	b = Board{
		{Mine, Closed, Mine},
		{Closed, Closed, Closed},
		{Mine, Closed, Mine},
	}
	assert.Equal(t, 0, b.CalcNeighbors(0, 0))
	assert.Equal(t, 2, b.CalcNeighbors(0, 1))
	assert.Equal(t, 2, b.CalcNeighbors(1, 0))
	assert.Equal(t, 4, b.CalcNeighbors(1, 1))
	assert.Equal(t, 0, b.CalcNeighbors(0, 2))
	assert.Equal(t, 0, b.CalcNeighbors(2, 0))

	// non-symmetric test
	b = Board{
		{Mine, Mine, Mine},
		{Closed, Mine, Mine},
		{Closed, Closed, Mine},
	}
	assert.Equal(t, 2, b.CalcNeighbors(0, 0))
	assert.Equal(t, 4, b.CalcNeighbors(0, 1))
	assert.Equal(t, 3, b.CalcNeighbors(1, 0))
	assert.Equal(t, 5, b.CalcNeighbors(1, 1))
	assert.Equal(t, 3, b.CalcNeighbors(0, 2))
	assert.Equal(t, 1, b.CalcNeighbors(2, 0))
	assert.Equal(t, 4, b.CalcNeighbors(1, 2))
	assert.Equal(t, 3, b.CalcNeighbors(2, 1))
	assert.Equal(t, 2, b.CalcNeighbors(2, 2))

	b = Board{
		{Closed, Closed, Closed},
		{Closed, Closed, Closed},
		{Closed, Closed, Closed},
	}
	for i, row := range b {
		for j := range row {
			if i == j {
				continue
			}
			assert.Equal(t, 0, b.CalcNeighbors(i, j))
		}
	}

	b = Board{
		{Mine, Mine, Mine},
		{Mine, Mine, Mine},
		{Mine, Mine, Mine},
	}
	assert.Equal(t, 3, b.CalcNeighbors(0, 0))
	assert.Equal(t, 5, b.CalcNeighbors(0, 1))
	assert.Equal(t, 5, b.CalcNeighbors(1, 0))
	assert.Equal(t, 8, b.CalcNeighbors(1, 1))
	assert.Equal(t, 3, b.CalcNeighbors(0, 2))
	assert.Equal(t, 3, b.CalcNeighbors(2, 0))
	assert.Equal(t, 5, b.CalcNeighbors(1, 2))
	assert.Equal(t, 5, b.CalcNeighbors(2, 1))
	assert.Equal(t, 3, b.CalcNeighbors(2, 2))
}

func TestOpenCell(t *testing.T) {
  cellIsMine := func(t *testing.T) {
    game := NewGame()
    game.board[3][3] = Mine
    game.OpenCell(3, 3)

    assert.Equal(t, Lose, game.state)
  }


  cellIsMine(t)
}
