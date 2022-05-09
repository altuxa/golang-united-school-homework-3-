package homework

import (
	"sort"
)

func reverse(input []int64) (result []int64) {
	// Place your code here
	slice := make([]int64, len(input))
	copy(slice, input)
	sort.Slice(slice, func(i, j int) bool { return slice[i] > slice[j] })
	return slice
}
