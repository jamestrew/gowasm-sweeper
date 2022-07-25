package main

import (
	"fmt"
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

type Board [][]Cell

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

type Game struct {
	width, height int
	mineCount     int
	board         *Board
	state         GameState
}

func (b Board) CalcNeighbors(y, x int) int {
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
	if g.board[x][y] == Mine {
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

func buildIntArray(width, height int) [][]Cell {
	ret := make([][]Cell, width)
	for i := range ret {
		ret[i] = make([]Cell, height)
	}
	return ret
}

// TODO: make random and also refactor this. this is dumb
func generateMines(width, height int) [][]bool {
	ret := buildBoolArray(width, height)
	for i := range ret {
		for j := range ret[i] {
			if i == j {
				ret[i][j] = true
			}
		}
	}

	return ret
}

func (b *Board) IntergrateMines(width, height, mines int) {
	mineBoard := generateMines(width, height)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if mineBoard[i][j] {
				(*b)[i][j] = Mine
			} else {
				(*b)[i][j] = Closed
			}
		}
	}
}

// TODO: get inputs from user
func GetCustomBoardParams() (int, int, int) {
	return 9, 9, 10
}

func GetBoardParams(difficulty DifficultyLevel) (int, int, int) {
	var width, height, mines int
	switch difficulty {
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

func NewBoard(width, height, mines int) *Board {
	board := make(Board, width)
	for i := range board {
		board[i] = make([]Cell, height)
	}

	board.IntergrateMines(width, height, mines)

	return &board
}

func NewGame(difficulty DifficultyLevel) *Game {
	width, height, mines := GetBoardParams(difficulty)
	board := NewBoard(width, height, mines)
	game := &Game{width, height, mines, board, Playing}
	return game
}

func NewTestGame(board *Board) *Game {
  b := *board
  width, height := len(b), len(b[0])
	game := &Game{width, height, -1, board, Playing}
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
