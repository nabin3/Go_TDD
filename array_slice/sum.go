package array_slice

func Sum(numbers []int) int {
	summation := 0
	for _, number := range numbers {
		summation += number
	}
	return summation
}

func SumAll(numSets ...[]int) []int {
	resultSet := make([]int, len(numSets))

	for i, set := range numSets {
		resultSet[i] = Sum(set)
	}

	return resultSet
}

func SumAllTails(numSets ...[]int) []int {
	resultSet := make([]int, len(numSets))

	for i, set := range numSets {
		if len(set) == 0 {
			resultSet[i] = 0
		} else {
			resultSet[i] = Sum(set[1:])
		}
	}

	return resultSet
}
