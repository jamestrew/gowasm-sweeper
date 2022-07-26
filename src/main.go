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

// TODO: deprecate above type for below struct
type Board struct {
	cells     [][]Cell
	width     int
	height    int
	mineCount int
}

type GameState uint8

const (
	Playing GameState = iota
	Win
	Lose
)

type DifficultyLevel uint8

const (
	Beginner DifficultyLevel = iota
	Intermediate
	Expert
	Custom
)

// TODO: make this apart of the Board struct (see above)
type Difficulty struct {
	width     int
	height    int
	mineCount int
}

type Game struct {
	difficulty Difficulty
	board      *Cells
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

// func (g *Game) OpenCell(x, y int) {
// 	if g.board[x][y] == Mine {
// 		g.state = Lose
// 	}
//
// 	// check neighboring cells and set the number for current cell
// 	// open any neighboring cell if current cell is a zero
// }

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

func mineCoordinates(difficulty Difficulty) [][2]int {
	mineCount := 0
	mines := coordinateArray(difficulty.mineCount, -1, -1)

	for mineCount < difficulty.mineCount {
		x := rand.Intn(difficulty.width)
		y := rand.Intn(difficulty.height)
		if !isDuplicateMineCoord(x, y, mines) {
			mines[mineCount] = [2]int{x, y}
			mineCount++
		}
	}

	return mines
}

// TODO: make random and also refactor this. this is dumb
func generateMines(difficulty Difficulty) [][]bool {
	ret := buildBoolArray(difficulty.width, difficulty.height)
	for i := range ret {
		for j := range ret[i] {
			if i == j {
				ret[i][j] = true
			}
		}
	}
	return ret
}

func (b *Cells) IntergrateMines(difficulty Difficulty) {
	mineCoords := mineCoordinates(difficulty)
	for _, coord := range mineCoords {
		x, y := coord[0], coord[1]
		(*b)[y][x] = Mine
	}
}

// TODO: get inputs from user
func GetCustomBoardParams() (int, int, int) {
	return 9, 9, 10
}

func GetBoardParams(level DifficultyLevel) Difficulty {
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
	return Difficulty{width, height, mines}
}

func CreateBlankBoard(difficulty Difficulty) *Cells {
	cells := make(Cells, difficulty.height)
	for i := range cells {
		cells[i] = make([]Cell, difficulty.width)
	}

	for _, row := range cells {
		for i := range row {
			row[i] = None
		}
	}
	return &cells
}

func NewBoard(difficulty Difficulty) *Cells {
	cells := CreateBlankBoard(difficulty)
	cells.IntergrateMines(difficulty)
	return cells
}

func NewGame(level DifficultyLevel) *Game {
	difficulty := GetBoardParams(level)
	board := NewBoard(difficulty)
	game := &Game{difficulty, board, Playing}
	return game
}

func NewTestGame(c *Cells) *Game {
	b := *c
	width, height := len(b), len(b[0])
	difficulty := Difficulty{width, height, -1}
	game := &Game{difficulty, c, Playing}
	return game
}

func (g Game) PrintBoard() {
	board := *(g.board)
	for i := range board {
		fmt.Println(board[i])
	}
}

func main() {
	fmt.Println("hello world!")
	game := NewGame(Beginner)
	game.PrintBoard()
	fmt.Println(game.board.CalcNeighbors(0, 0))

  b := Board{
    {Mine, Closed, Mine},
    {Closed, Closed, Closed},
    {Mine, Closed, Mine},
  }

  testGame := NewTestGame(&b)
  testGame.PrintBoard()
}
