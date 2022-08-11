package main

import (
	"syscall/js"

	g "github.com/jamestrew/gowasm-sweeper/pkg/game"
)

var game *g.Game

func newGame() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		difficulty := g.DifficultyLevel(args[0].Int())
		g.CustomWidth = args[1].Int()
		g.CustomHeight = args[2].Int()
		g.CustomMineCount = args[3].Int()
		game, _ = g.NewGame(difficulty)
		game.FillMines()
		return game.AsJson()
	})
}

func openCell() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		x, y := args[0].Int(), args[1].Int()
		game.OpenCell(x, y)
		return game.AsJson()
	})
}

func main() {
	ch := make(chan struct{}, 0)
	js.Global().Set("newGame", newGame())
	js.Global().Set("openCell", openCell())
	<-ch
}
