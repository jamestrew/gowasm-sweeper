package main

import (
	"fmt"
)

type Cell int8

const (
	Mine Cell = iota
	Closed
	Opened
	Flagged
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
)

type Board [][]Cell

type GameState uint8

const (
	Playing GameState = 1
	Win               = 2
	Lose              = 3
)

// placeholder for adding set difficulties
type Difficulty uint8

const (
	Beginner     Difficulty = 10 // 9x9, 10 mines
	Intermediate            = 40 // 16x16, 40 mines
	Expert                  = 99 // 30x16, 99 mines
)

type Game struct {
	width, height int
	mineCount     int // 1, 2, 3 = beginner, intermediate, expert
	board         Board
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

func (g *Game) IntegrateMines() {
	mines := generateMines(g.width, g.height)
	for i := 0; i < g.width; i++ {
		for j := 0; j < g.height; j++ {
			if mines[i][j] {
        fmt.Println("mining")
				g.board[i][j] = Mine
			} else {
        g.board[i][j] = Closed
      }
		}
	}

}

func NewGame() *Game {
	width, height := 9, 9
	board := buildIntArray(width, height)
	game := &Game{width, height, 10, board, Playing}
	game.IntegrateMines()
	return game
}

func (g *Game) PrintBoard() {
	for i := range g.board {
		fmt.Println(g.board[i])
	}
}

func main() {
	fmt.Println("hello world!")
	game := NewGame()
	game.PrintBoard()
	fmt.Println(game.board.CalcNeighbors(0, 0))

	fmt.Println(generateMines(9, 9))
}
