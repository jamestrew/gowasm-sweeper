package main

import (
	"fmt"
)

// board:
// 0-8 denotes neighboring mines
// 9 denotes mine
// -1 denotes unopened
// -2 denote flag
type Board [][]int

type Game struct {
	width, height int
	difficulty    int // 1, 2, 3 = beginner, intermediate, expert
	board         Board
}


func buildBoolArray(width, height int) [][]bool {
	ret := make([][]bool, width)
	for i := range ret {
		ret[i] = make([]bool, height)
	}
	return ret
}

func buildIntArray(width, height int) [][]int {
	ret := make([][]int, width)
	for i := range ret {
		ret[i] = make([]int, height)
	}
	return ret
}

// TODO: make random
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
				g.board[i][j] = 9
			}
		}
	}

}

func NewGame() *Game {
	width, height := 9, 9
	board := buildIntArray(width, height)
	game := &Game{width, height, 1, board}
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
	board := NewGame()
	board.PrintBoard()
}
