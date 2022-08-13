package main

import (
	"syscall/js"

	g "github.com/jamestrew/gowasm-sweeper/pkg/game"
)

var game *g.Game

func newGame(this js.Value, args []js.Value) interface{} {
	difficulty := g.DifficultyLevel(args[0].Int())
	g.CustomWidth = args[1].Int()
	g.CustomHeight = args[2].Int()
	g.CustomMineCount = args[3].Int()
	game, _ = g.NewGame(difficulty)
	// TODO: create random seed
	game.FillMines()
	// game.CalcAllNeighbors()
	return game.AsJson()
}

func openCell(this js.Value, args []js.Value) interface{} {
	x, y := args[0].Int(), args[1].Int()
	game.OpenCell(x, y)
	return game.AsJson()
}

func main() {
	js.Global().Set("newGame", js.FuncOf(newGame))
	js.Global().Set("openCell", js.FuncOf(openCell))
	<-make(chan bool)
}
