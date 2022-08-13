package main

import (
	"syscall/js"

	g "github.com/jamestrew/gowasm-sweeper/pkg/game"
)

var game *g.Game

func newGame(this js.Value, args []js.Value) interface{} {
	gameParams := args[0]
	difficulty := g.DifficultyLevel(gameParams.Get("difficulty").Int())
	if difficulty == g.Custom {
		g.CustomWidth = gameParams.Get("width").Int()
		g.CustomHeight = gameParams.Get("height").Int()
		g.CustomMineCount = gameParams.Get("mineCount").Int()
	}
	game, _ = g.NewGame(difficulty)
	// TODO: create random seed
	game.FillMines()
	game.CalcAllNeighbors()
	return game.AsJson()
}

func openCell(this js.Value, args []js.Value) interface{} {
	x, y := args[0].Int(), args[1].Int()
	game.OpenCell(x, y)
	return game.AsJson()
}

func getState(this js.Value, args []js.Value) interface{} {
	return game.AsJson()
}

func main() {
	js.Global().Set("newGame", js.FuncOf(newGame))
	js.Global().Set("openCell", js.FuncOf(openCell))
	js.Global().Set("getState", js.FuncOf(getState))
	<-make(chan bool)
}
