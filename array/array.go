package array

func Add(numbers [5]int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}
