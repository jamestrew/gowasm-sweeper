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

type Difficulty struct {
	width     int
	height    int
	mineCount int
}

type Game struct {
	difficulty    Difficulty
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

func (b *Board) IntergrateMines(difficulty Difficulty) {
	mineBoard := generateMines(difficulty)
	for i := 0; i < difficulty.width; i++ {
		for j := 0; j < difficulty.height; j++ {
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

func NewBoard(difficulty Difficulty) *Board {
	board := make(Board, difficulty.width)
	for i := range board {
		board[i] = make([]Cell, difficulty.height)
	}

	board.IntergrateMines(difficulty)

	return &board
}

func NewGame(level DifficultyLevel) *Game {
	difficulty := GetBoardParams(level)
	board := NewBoard(difficulty)
	game := &Game{difficulty, board, Playing}
	return game
}

func NewTestGame(board *Board) *Game {
	b := *board
	width, height := len(b), len(b[0])
  difficulty := Difficulty{width, height, -1}
	game := &Game{difficulty, board, Playing}
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
