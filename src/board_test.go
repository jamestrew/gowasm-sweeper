package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

  // assert equality
  assert.Equal(t, 123, 123, "they should be equal")

  // assert inequality
  assert.NotEqual(t, 123, 456, "they should not be equal")

}

func TestNewGame(t *testing.T) {
  game := NewGame()

  assert.Equal(t, game.width, 9)
  assert.Equal(t, game.height, 9)
  assert.Equal(t, game.width, 9)
  assert.Equal(t, game.state, Playing)
  assert.Equal(t, game.state, Playing)

  mineCount := 0
  for _, row := range game.board {
    for _, cell := range row {
      if cell == Mine {
        mineCount++
      }
    }
  }
  assert.Equal(t, mineCount, 10)
}
