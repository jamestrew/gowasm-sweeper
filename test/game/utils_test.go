package game_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	g "github.com/jamestrew/gowasm-sweeper/pkg/game"
)

func TestPosArray(t *testing.T) {
	var posArray = g.ExportPosArray
	expected := []g.Pos{
		{X: -1, Y: -1},
		{X: -1, Y: -1},
		{X: -1, Y: -1},
		{X: -1, Y: -1},
	}
	assert.Equal(t, expected, posArray(4))
}

func TestIsDuplicateMinePos(t *testing.T) {
	var isDuplicateMinePos = g.ExportIsDuplicateMinePos
	mines := []g.Pos{{X: 0, Y: 0}}

	assert.Equal(t, true, isDuplicateMinePos(g.Pos{X: 0, Y: 0}, mines))
	assert.Equal(t, false, isDuplicateMinePos(g.Pos{X: 1, Y: 1}, mines))
}

func TestCreateMinePositions(t *testing.T) {
	var createMinePositions = g.ExportCreateMinePositions

	rand.Seed(69420)
	g.CustomWidth = 4
	g.CustomHeight = 3
	g.CustomMineCount = 4
	game, _ := g.NewGame(g.Custom)

	expected := []g.Pos{
		{X: 0, Y: 1},
		{X: 3, Y: 0},
		{X: 1, Y: 1},
		{X: 2, Y: 2},
	}
	assert.Equal(t, expected, createMinePositions(game))
}

func TestCellNeighbors(t *testing.T) {
	var cellNeighbors = g.ExportCellNeighbors

	g.CustomWidth = 4
	g.CustomHeight = 3
	g.CustomMineCount = 4
	game, _ := g.NewGame(g.Custom)
	/*
		{0, 1, 9, 1},
		{0, 1, 1, 1},
		{0, 1, 9, 1},
	*/

	inTheMiddle := func() {
		expected := []g.Pos{
			{X: 1, Y: 0},
			{X: 2, Y: 0},
			{X: 3, Y: 0},
			{X: 1, Y: 1},
			{X: 3, Y: 1},
			{X: 1, Y: 2},
			{X: 2, Y: 2},
			{X: 3, Y: 2},
		}
		assert.ElementsMatch(t, expected, cellNeighbors(game, 2, 1))
	}

	topEdge := func() {
		expected := []g.Pos{
			{X: 1, Y: 0},
			{X: 3, Y: 0},
			{X: 1, Y: 1},
			{X: 2, Y: 1},
			{X: 3, Y: 1},
		}
		assert.ElementsMatch(t, expected, cellNeighbors(game, 2, 0))
	}

	corner := func() {
		expected := []g.Pos{
			{X: 1, Y: 0},
			{X: 0, Y: 1},
			{X: 1, Y: 1},
		}
		assert.ElementsMatch(t, expected, cellNeighbors(game, 0, 0))
	}

	inTheMiddle()
	topEdge()
	corner()
}
