package calculator

func Add(numbers []int) int {
	sum := 0

	for _, n := range numbers {
		sum = sum + n
	}

	return sum
}
