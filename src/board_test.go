package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestNewGame(t *testing.T) {
// 	game := NewGame(Beginner)
//
// 	assert.Equal(t, game.difficulty.width, 9)
// 	assert.Equal(t, game.difficulty.height, 9)
// 	assert.Equal(t, game.state, Playing)
// 	assert.Equal(t, game.state, Playing)
//
// 	mineCount := 0
// 	for _, row := range *(game.board) {
// 		for _, cell := range row {
// 			if cell == Mine {
// 				mineCount++
// 			}
// 		}
// 	}
// 	assert.Equal(t, 9, mineCount)
// }

func TestCalcNeighbors(t *testing.T) {

	b := Cells{
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

	b = Cells{
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

	b = Cells{
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
	b = Cells{
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

	b = Cells{
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

	b = Cells{
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

func TestCoordinateArray(t *testing.T) {
	zeroCount := func() {
		ret := coordinateArray(0, -1, -1)
		assert.Equal(t, ret, [][2]int{})
	}

	nonZeroCount := func() {
		ret := coordinateArray(3, -1, -1)
		assert.Equal(t, ret, [][2]int{{-1, -1}, {-1, -1}, {-1, -1}})
	}

	zeroCount()
	nonZeroCount()
}

func TestIsDuplicateMineCoord(t *testing.T) {
	emptyMines := func() {
		mines := [][2]int{{-1, -1}, {-1, -1}, {-1, -1}}
		assert.False(t, isDuplicateMineCoord(0, 0, mines))
	}

	nonDuplicate := func() {
		mines := [][2]int{{1, 1}}
		assert.False(t, isDuplicateMineCoord(0, 0, mines))
	}

	duplicate := func() {
		mines := [][2]int{{1, 1}}
		assert.True(t, isDuplicateMineCoord(1, 1, mines))
	}

	emptyMines()
	nonDuplicate()
	duplicate()
}

// this test is pretty dodgy
func TestMineCoordinates(t *testing.T) {
	rand.Seed(69421)
	difficulty := Difficulty{3, 4, 3}
	assert.Equal(t, mineCoordinates(difficulty), [][2]int{{2, 2}, {0, 2}, {2, 3}})
	assert.Equal(t, mineCoordinates(difficulty), [][2]int{{1, 3}, {1, 2}, {0, 0}})
	assert.Equal(t, mineCoordinates(difficulty), [][2]int{{2, 1}, {0, 1}, {1, 2}})
	assert.Equal(t, mineCoordinates(difficulty), [][2]int{{1, 1}, {0, 1}, {2, 0}})
}

func TestIntegrateMines(t *testing.T) {
	rand.Seed(69421)
	b := &Cells{
		{None, None, None},
		{None, None, None},
		{None, None, None},
		{None, None, None},
	}
	difficulty := Difficulty{3, 4, 3}
	b.IntergrateMines(difficulty)

	assert.Equal(t, *b, Cells{
		{None, None, None},
		{None, None, None},
		{Mine, None, Mine},
		{None, None, Mine},
	})
}

func TestCreateBlankBoard(t *testing.T) {
	difficulty := Difficulty{3, 4, 3}
	board := *CreateBlankBoard(difficulty)
	assert.Equal(t, board, Cells{
		{None, None, None},
		{None, None, None},
		{None, None, None},
		{None, None, None},
	})
}
