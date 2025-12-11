package combinations

import (
	"gonum.org/v1/gonum/stat/combin"
)

func Choose[T any](n int, from []T) [][]T {
	indices := combin.Combinations(len(from), n)
	return mapIndices(from, indices)
}

func mapIndices[T any](from []T, indices [][]int) [][]T {
	result := make([][]T, len(indices))
	for i, list := range indices {
		result[i] = make([]T, len(list))
		for j, index := range list {
			result[i][j] = from[index]
		}
	}
	return result
}
