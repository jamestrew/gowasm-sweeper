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

func (g *Game) cellNeighbors(x, y int) []Pos {
	dxs := [3]int{-1, 0, 1}
	dys := [3]int{0, 1, -1}

	isOutOfBounds := func(dx, dy int) bool {
		return x+dx >= g.Width || x+dx < 0 || y+dy >= g.Height || y+dy < 0
	}

	positions := []Pos{}
	for _, dx := range dxs {
		for _, dy := range dys {
			if isOutOfBounds(dx, dy) || (dx == 0 && dy == 0) {
				continue
			}
			positions = append(positions, Pos{x + dx, y + dy})
		}
	}
	return positions
}
