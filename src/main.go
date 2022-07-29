package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Cell int8

const (
	None Cell = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Mine
	Closed
	Flagged
)


type Cells [][]Cell
type BoardParams struct {
	width     int
	height    int
	mineCount int
}

type Board struct {
	cells  Cells
	params BoardParams
}

type GameState uint8

const (
	Playing GameState = iota
	Win
	Lose
)

type Difficulty uint8

const (
	Beginner Difficulty = iota
	Intermediate
	Expert
	Custom
)

type Game struct {
	difficulty Difficulty
	board      Board
	state      GameState
}

func (b Cells) CalcNeighbors(y, x int) int {
	count := 0
	dxs := []int{-1, 0, 1}
	dys := []int{0, 1, -1}

	isOutOfBounds := func(dx, dy int) bool {
		return x+dx > 2 || x+dx < 0 || y+dy > 2 || y+dy < 0
	}

	for _, dx := range dxs {
		for _, dy := range dys {
			if isOutOfBounds(dx, dy) || (dx == 0 && dy == 0) {
				continue
			}
			if b[y+dy][x+dx] == Mine {
				count++
			}
		}
	}
	return count
}

func (g *Game) OpenCell(x, y int) {
	if g.board.cells[x][y] == Mine {
		g.state = Lose
	}

	// check neighboring cells and set the number for current cell
	// open any neighboring cell if current cell is a zero
}

func buildBoolArray(width, height int) [][]bool {
	ret := make([][]bool, width)
	for i := range ret {
		ret[i] = make([]bool, height)
	}
	return ret
}

func isDuplicateMineCoord(newX, newY int, mines [][2]int) bool {
	for _, coord := range mines {
		x, y := coord[0], coord[1]
		if x == newX && y == newY {
			return true
		}
	}
	return false

}

func coordinateArray(count, defaultX, defaultY int) [][2]int {
	ret := make([][2]int, count)
	for i := 0; i < count; i++ {
		ret[i] = [2]int{defaultX, defaultY}
	}
	return ret
}

func mineCoordinates(boardParams BoardParams) [][2]int {
	mineCount := 0
	mines := coordinateArray(boardParams.mineCount, -1, -1)

	for mineCount < boardParams.mineCount {
		x := rand.Intn(boardParams.width)
		y := rand.Intn(boardParams.height)
		if !isDuplicateMineCoord(x, y, mines) {
			mines[mineCount] = [2]int{x, y}
			mineCount++
		}
	}

	return mines
}

func (b Board) IntergrateMines() {
	mineCoords := mineCoordinates(b.params)
	for _, coord := range mineCoords {
		x, y := coord[0], coord[1]
		(b.cells)[y][x] = Mine
	}
}

// TODO: get inputs from user
func GetCustomBoardParams() (int, int, int) {
	return 9, 9, 10
}

func GetBoardParams(level Difficulty) BoardParams {
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
	return BoardParams{width, height, mines}
}

func CreateBlankBoard(difficulty Difficulty) Board {
	params := GetBoardParams(difficulty)
	cells := make(Cells, params.height)
	for i := range cells {
		cells[i] = make([]Cell, params.width)
	}

	for _, row := range cells {
		for i := range row {
			row[i] = None
		}
	}
	return Board{cells, params}
}

func NewBoard(difficulty Difficulty) Board {
	cells := CreateBlankBoard(difficulty)
	cells.IntergrateMines()
	return cells
}

func NewGame(difficulty Difficulty) *Game {
	board := NewBoard(difficulty)
	game := &Game{difficulty, board, Playing}
	return game
}

func NewTestGame(c Cells) *Game {
	height, width := len(c), len(c[0])
	board := Board{c, BoardParams{width, height, -1}}
	game := &Game{Custom, board, Playing}
	return game
}

func (g Game) PrintBoard() {
	for i := range g.board.cells {
		fmt.Println(g.board.cells[i])
	}
}

func main() {
	fmt.Println("hello world!")
	rand.Seed(time.Now().UnixNano())
	game := NewGame(Beginner)
	game.PrintBoard()
}
