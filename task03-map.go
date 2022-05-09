package homework

func sortMapValues(input map[int]string) (result []string) {
	// Place your code here
	var res []string
	for i := 0; i < len(input); i++ {
		res = append(res, input[i])
	}
	return res
}
