package utils


func InitMatrix[T any](width, height int) [][]T {
  ret := make([][]T, height)
  for i := range ret {
    ret[i] = make([]T, width)
  }
  return ret
}

func FillMatrix[T any](matrix [][]T, fillValue T) [][]T {
	for _, row := range matrix {
		for i := range row {
			row[i] = fillValue
		}
	}
  return matrix
}

func InitFilledMatrix[T any](width, height int, fillValue T) [][]T {
  matrix := InitMatrix[T](width, height)
  return FillMatrix(matrix, fillValue)
}
