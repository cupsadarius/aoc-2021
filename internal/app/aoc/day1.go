package aoc

func SolveDay1Part1(input []int) int {
	prevValue := 0
	increases := -1
	for _, depth := range input {
		if depth > prevValue {
			increases += 1
		}
		prevValue = depth
	}

	return increases
}

func SolveDay1Part2(input []int) int {
	sums := []int{}
	for index :=0; index < len(input) - 2; index ++ {
		tempSum := 0
		for i :=0; i < 3; i++ {
			tempSum += input[index + i]
		}
		sums = append(sums, tempSum)
	}
	return SolveDay1Part1(sums)
}