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

type GameState uint8

const (
	Unstarted GameState = iota
	Playing
	Win
	Lose
)

type Game struct {
	Width      int
	Height     int
	MineCount  int
	Difficulty DifficultyLevel
	Mines      [][]int
	Open       [][]bool
	Flagged    [][]bool
	State      GameState
}

type Pos struct {
	X, Y int
}

var CustomWidth int = 4
var CustomHeight int = 3
var CustomMineCount int = 3

func NewGame(difficulty DifficultyLevel) (*Game, error) {
	width, height, mineCount := GetBoardParams(difficulty)
	if mineCount > width*height {
		errMsg := fmt.Sprintf("Mine count (%d) too large for given board dimensions (%dx%d)", mineCount, width, height)
		return nil, errors.New(errMsg)
	}

	mines := utils.InitBlankMatrix[int](width, height)
	open := utils.InitBlankMatrix[bool](width, height)
	flagged := utils.InitBlankMatrix[bool](width, height)
	return &Game{width, height, mineCount, difficulty, mines, open, flagged, Playing}, nil
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
	count := 0
	for _, pos := range g.cellNeighbors(x, y) {
		if g.Mines[pos.Y][pos.X] == 9 {
			count++
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

func (g *Game) CountBlankNeighbors(x, y int) int {
	count := 0
	for _, pos := range g.cellNeighbors(x, y) {
		if g.Mines[pos.Y][pos.X] == 0 {
			count++
		}
	}
	return count
}

func (g *Game) OpenBlankCells(x, y int) {
	if g.CountBlankNeighbors(x, y) == 0 || g.Open[y][x] {
		return
	}

	g.Open[y][x] = true
	for _, pos := range g.cellNeighbors(x, y) {
		if g.Mines[pos.Y][pos.X] == 0 {
			g.OpenBlankCells(pos.X, pos.Y)
		}
	}
}
