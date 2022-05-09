package homework

func average(input [15]float32) (result float32) {
	// Place your code here
	var sum float32
	var ln float32
	for _, i := range input {
		if i != 0 {
			sum = sum + i
			ln++
		}
	}
	return (sum / ln)
}
