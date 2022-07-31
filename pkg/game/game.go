package game

import (
	"errors"
	"fmt"

	"github.com/jamestrew/gowasm-sweeper/pkg/utils"
)

type DifficultyLevel uint8

const (
	Beginner DifficultyLevel = iota
	Intermediate
	Expert
	Custom
)

type Game struct {
	Width      int
	Height     int
	MineCount  int
	Difficulty DifficultyLevel
	Mines      [][]int
	Open       [][]bool
	Flagged    [][]bool
}

type Pos struct {
	X, Y int
}

var CustomWidth int = 4
var CustomHeight int = 3
var CustomMineCount int = 3

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

func NewGame(difficulty DifficultyLevel) (*Game, error) {
	width, height, mineCount := GetBoardParams(difficulty)
	if mineCount > width*height {
		errMsg := fmt.Sprintf("Mine count (%d) too large for given board dimensions (%dx%d)", mineCount, width, height)
		return nil, errors.New(errMsg)
	}

	mines := utils.InitBlankMatrix[int](width, height)
	open := utils.InitBlankMatrix[bool](width, height)
	flagged := utils.InitBlankMatrix[bool](width, height)
	return &Game{width, height, mineCount, difficulty, mines, open, flagged}, nil
}

func (g *Game) CalcAllNeighbors() {
	for i, row := range g.Mines {
		for j, cell := range row {
			if cell != 9 {
				row[j] = g.CalcCellNeighbors(j, i)
			}
		}
	}

}

func (g *Game) CalcCellNeighbors(x, y int) int {
	dxs := [3]int{-1, 0, 1}
	dys := [3]int{0, 1, -1}

	isOutOfBounds := func(dx, dy int) bool {
		return x+dx >= g.Width || x+dx < 0 || y+dy >= g.Height || y+dy < 0
	}

	count := 0
	for _, dx := range dxs {
		for _, dy := range dys {
			if isOutOfBounds(dx, dy) || (dx == 0 && dy == 0) {
				continue
			}
			if g.Mines[y+dy][x+dx] == 9 {
				count++
			}
		}
	}

	return count
}

func (g *Game) FillMines() {
	minePositions := createMinePositions(g)
	for _, pos := range minePositions {
		g.Mines[pos.Y][pos.X] = 9 // itself a mine
	}
}
