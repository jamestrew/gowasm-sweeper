package game

import (
	"math/rand"
)

func posArray(mineCount int) []Pos {
	positions := make([]Pos, mineCount)
	for i := range positions {
		positions[i] = Pos{-1, -1}
	}
	return positions
}

func isDuplicateMinePos(testPosition Pos, minePositions []Pos) bool {
	for _, pos := range minePositions {
		if testPosition.X == pos.X && testPosition.Y == pos.Y {
			return true
		}
	}
	return false
}

func createMinePositions(game *Game) []Pos {
	minesCreated := 0
	mines := posArray(game.MineCount)

	for minesCreated < game.MineCount {
		testPosition := Pos{rand.Intn(game.Width), rand.Intn(game.Height)}
		if !isDuplicateMinePos(testPosition, mines) {
			mines[minesCreated] = testPosition
      minesCreated++
		}
	}
	return mines
}
