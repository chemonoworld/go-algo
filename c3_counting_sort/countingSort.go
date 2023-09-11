package c3_counting_sort

func CountingSort(input []int, lowerLimit, upperLimit int) []int {
	countSize := upperLimit - lowerLimit + 1
	offset := lowerLimit

	count := make([]int, countSize)

	for i := 0; i < len(input); i++ {
		count[input[i]-offset]++
	}

	var result []int
	for i := 0; i < countSize; i++ {
		for j := 0; j < count[i]; j++ {
			result = append(result, i+offset)
		}
	}

	return result
}
