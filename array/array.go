package array

func Add(numbers []int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}

func SumSlices(slices ...[]int) []int {
	var result []int

	for _, v := range slices {
		result = append(result, Add(v))
	}

	return result
}
