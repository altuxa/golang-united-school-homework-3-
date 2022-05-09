package homework

func reverse(input []int64) (result []int64) {
	// Place your code here
	slice := make([]int64, len(input))
	copy(slice, input)
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
