package convertors

import (
	"aoc/internal/app/aoc/types"
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

func GetAsBoards(input []string) ([]int, []types.Bingo) {
	sequence := []int{}
	for _, entry := range strings.Split(input[0], ",") {
		value, _ := strconv.Atoi(entry)
		sequence = append(sequence, value)
	}

	boards := []types.Bingo{}
	slice := input[1:]
	for index := 0; index < len(slice)-5; index += 6 {
		board := types.Bingo{
			Board:     make([][]types.BoardCell, 5),
			Completed: false,
		}
		for i := 1; i <= 5; i++ {
			for _, entry := range strings.Split(slice[index+i], " ") {
				if entry == "" {
					continue
				}
				value, _ := strconv.Atoi(entry)
				board.Board[i-1] = append(board.Board[i-1], types.BoardCell{Value: value, Marked: false})
			}
		}

		boards = append(boards, board)
	}

	return sequence, boards
}

func GetAsLines(input []string) []types.Line {
	lines := []types.Line{}
	for _, entry := range input {
		words := strings.Split(entry, " -> ")
		points := make([]types.Point, 2)
		for index, word := range words {
			coords := strings.Split(word, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			points[index] = types.Point{X: x, Y: y, Value: 0}
		}
		lines = append(lines, types.Line{A: points[0], B: points[1]})
	}

	return lines
}
