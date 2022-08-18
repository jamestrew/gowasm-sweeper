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

func isBadMinePos(testPosition Pos, minePositions []Pos) bool {
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
		if !isBadMinePos(testPosition, mines) {
			mines[minesCreated] = testPosition
			minesCreated++
		}
	}
	return mines
}

func (g *Game) cleanMinePositions(noGoZone []Pos) []Pos {
	minesCreated := 0
	mines := posArray(g.MineCount)

	for minesCreated < g.MineCount {
		pos := Pos{rand.Intn(g.Width), rand.Intn(g.Height)}
		if !isBadMinePos(pos, mines) && !isBadMinePos(pos, noGoZone) {
			mines[minesCreated] = pos
			minesCreated++
		}
	}
	return mines
}

func (g *Game) CellNeighbors(x, y int) []Pos {
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

func (g *Game) positionsMineCount(positions []Pos) int {
	mineCount := 0
	for _, pos := range positions {
		if g.Mines[pos.Y][pos.X] == 9 {
			g.Mines[pos.Y][pos.X] = 0
			mineCount++
		}
	}
	return mineCount
}

func (g *Game) randomCellWith(cellType int) Pos {
	for {
		pos := Pos{rand.Intn(g.Width), rand.Intn(g.Height)}
		if g.Mines[pos.Y][pos.X] == cellType {
			return pos
		}
	}
}

func (g *Game) randomCellNot(cellType int) Pos {
	for {
		pos := Pos{rand.Intn(g.Width), rand.Intn(g.Height)}
		if g.Mines[pos.Y][pos.X] != cellType {
			return pos
		}
	}
}

// TODO
func GetCustomBoardParams() (int, int, int) {
	return CustomWidth, CustomHeight, CustomMineCount
}

func GetBoardParams(level DifficultyLevel) (int, int, int) {
	var width, height, mines int
	switch level {
	case Beginner:
		width, height, mines = 9, 9, 10
	case Intermediate:
		width, height, mines = 16, 16, 40
	case Expert:
		width, height, mines = 30, 16, 99
	case Custom:
		width, height, mines = GetCustomBoardParams()
	}
	return width, height, mines
}

func (g *Game) flagCountMatchesCell(x, y int) bool {
	flagCount := 0
	for _, pos := range g.CellNeighbors(x, y) {
		if g.Flagged[pos.Y][pos.X] {
			flagCount++
		}
	}
	return flagCount == g.Mines[y][x]
}

func (g *Game) playable() bool {
	return g.State == Unstarted || g.State == Playing
}

func (g *Game) isSpaciousBoard() bool {
	return (g.Width*g.Height)-g.MineCount >= 9
}
