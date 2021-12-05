package convertors

import (
	types "aoc/internal/app/aoc/types"
	"strconv"
	"strings"
)

func GetStringListAsInt(input []string) []int {
	values := []int{}
	for _, entry := range input {
		value, _ := strconv.Atoi(entry)
		values = append(values, value)
	}
	return values
}

func GetAsHeading(input []string) []types.Heading {
	actions := []types.Heading{}
	for _, entry := range input {
		words := strings.Fields(entry)
		change, _ := strconv.Atoi(words[1])
		action := types.Heading{
			Direction: words[0],
			Change:    change,
		}
		actions = append(actions, action)
	}
	return actions
}

func GetAsBytes(input []string) [][]int {
	bytes := [][]int{}
	for _, entry := range input {
		digits := []int{}
		for _, char := range strings.Split(entry, "") {
			digit, _ := strconv.Atoi(char)
			digits = append(digits, digit)
		}
		bytes = append(bytes, digits)
	}
	return bytes
}
