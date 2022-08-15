package main

import (
	"fmt"
	"math/rand"
	"syscall/js"
	"time"

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
	rand.Seed(time.Now().UnixNano())
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

func flagCell(this js.Value, args []js.Value) interface{} {
	x, y := args[0].Int(), args[1].Int()
	game.ToggleFlag(x, y)
	return game.AsJson()
}

func chordedOpen(this js.Value, args []js.Value) interface{} {
	x, y := args[0].Int(), args[1].Int()
	game.ChordedOpen(x, y)
	return game.AsJson()
}

func debugMines(this js.Value, args []js.Value) interface{} {
	fmt.Println(game.Mines)
	return nil
}

func debugState(this js.Value, args []js.Value) interface{} {
	return game.AsJson()
}

func main() {
	js.Global().Set("newGame", js.FuncOf(newGame))
	js.Global().Set("openCell", js.FuncOf(openCell))
	js.Global().Set("flagCell", js.FuncOf(flagCell))
	js.Global().Set("chordedOpen", js.FuncOf(chordedOpen))
	js.Global().Set("debugMines", js.FuncOf(debugMines))
	js.Global().Set("debugState", js.FuncOf(debugState))
	<-make(chan bool)
}
