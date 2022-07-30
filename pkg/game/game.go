package game

import (
	"github.com/jamestrew/gowasm-sweeper/pkg/utils"
)

type DifficultyLevel uint8

const (
	Beginner DifficultyLevel = iota
	Intermediate
	Expert
	Custom
)

type Game struct {
	Width      int
	Height     int
	MineCount  int
	Difficulty DifficultyLevel
	mines      [][]int
	open       [][]bool
	flagged    [][]bool
}

// TODO
func GetCustomBoardParams() (int, int, int) {
	return 9, 9, 10
}

func GetBoardParams(level DifficultyLevel) (int, int, int) {
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
	return width, height, mines
}

func NewGame(difficulty DifficultyLevel) *Game {
	width, height, mineCount := GetBoardParams(difficulty)
	mines := utils.InitMatrix[int](width, height)
	open := utils.InitMatrix[bool](width, height)
	flagged := utils.InitMatrix[bool](width, height)
	return &Game{width, height, mineCount, difficulty, mines, open, flagged}
}
