package utils_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	u "github.com/jamestrew/gowasm-sweeper/pkg/utils"
)

func TestInitMatrix(t *testing.T) {
	intMatrix := func() {
		expected := [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		}
		assert.Equal(t, expected, u.InitBlankMatrix[int](3, 4))
	}

	boolMatrix := func() {
		expected := [][]bool{
			{false, false, false},
			{false, false, false},
			{false, false, false},
			{false, false, false},
		}
		assert.Equal(t, expected, u.InitBlankMatrix[bool](3, 4))
	}

	intMatrix()
	boolMatrix()
}

func TestFillMatrix(t *testing.T) {
	intMatrix := func() {
		matrix := [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		}
		expected := [][]int{
			{42, 42, 42},
			{42, 42, 42},
			{42, 42, 42},
			{42, 42, 42},
		}
		assert.Equal(t, expected, u.FillMatrix(matrix, 42))
	}

	boolMatrix := func() {
		matrix := [][]bool{
			{false, false, false},
			{false, false, false},
			{false, false, false},
			{false, false, false},
		}
		expected := [][]bool{
			{true, true, true},
			{true, true, true},
			{true, true, true},
			{true, true, true},
		}
		assert.Equal(t, expected, u.FillMatrix(matrix, true))
	}

	intMatrix()
	boolMatrix()
}

func TestInitFilledMatrix(t *testing.T) {
	intMatrix := func() {
		expected := [][]int{
			{42, 42, 42},
			{42, 42, 42},
			{42, 42, 42},
			{42, 42, 42},
		}
		assert.Equal(t, expected, u.InitFilledMatrix(3, 4, 42))
	}

	boolMatrix := func() {
		expected := [][]bool{
			{true, true, true},
			{true, true, true},
			{true, true, true},
			{true, true, true},
		}
		assert.Equal(t, expected, u.InitFilledMatrix(3, 4, true))
	}

	intMatrix()
	boolMatrix()
}
