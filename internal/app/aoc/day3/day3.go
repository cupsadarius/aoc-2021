package day3

import (
	"strconv"
)

func Part1(input [][]int) int64 {
	sum := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, word := range input {
		for i, v := range word {
			sum[i] += v
		}
	}
	gammaBytes := ""
	half := len(input) / 2
	for _, v := range sum {
		if v > half {
			gammaBytes += "1"
		} else {
			gammaBytes += "0"
		}
	}

	gamma, _ := strconv.ParseInt(gammaBytes, 2, 64)
	epsilon := 0xfff - gamma
	return gamma * epsilon
}

func Part2(input [][]int) int64 {
	wordLength := len(input[0])
	mostCommon := input
	leastCommon := input
	for position := 0; position < wordLength; position++ {
		if len(mostCommon) > 1 {
			mostCommon = filterByMostCommonOnPosition(mostCommon, position)
		}
		if len(leastCommon) > 1 {
			leastCommon = filterByLeastCommonOnPosition(leastCommon, position)
		}
	}

	return byteToInt(mostCommon[0]) * byteToInt(leastCommon[0])
}

func filterByMostCommonOnPosition(input [][]int, position int) [][]int {
	filtered := [][]int{}
	mostCommon := getMostCommon(input, position)
	for _, word := range input {
		if word[position] == mostCommon {
			filtered = append(filtered, word)
		}
	}

	return filtered
}

func filterByLeastCommonOnPosition(input [][]int, position int) [][]int {
	filtered := [][]int{}
	leastCommon := getLeastCommon(input, position)
	for _, word := range input {
		if word[position] == leastCommon {
			filtered = append(filtered, word)
		}
	}

	return filtered
}

func getMostCommon(input [][]int, position int) int {
	half := float64(len(input)) / float64(2)
	sum := 0
	for _, word := range input {
		sum += word[position]
	}
	if float64(sum) > half || float64(sum) == half {
		return 1
	} else {
		return 0
	}
}
func getLeastCommon(input [][]int, position int) int {
	half := float64(len(input)) / float64(2)
	sum := 0
	for _, word := range input {
		sum += word[position]
	}
	if float64(sum) > half || float64(sum) == half {
		return 0
	} else {
		return 1
	}
}

func byteToInt(input []int) int64 {
	bytes := ""
	for _, v := range input {
		bytes += strconv.Itoa(v)
	}

	value, _ := strconv.ParseInt(bytes, 2, 64)

	return int64(value)
}
