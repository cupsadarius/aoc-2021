package day4

import (
	"aoc/internal/app/aoc/types"
)

func Part1(sequence []int, boards []types.Bingo) int {
	for _, number := range sequence {
		for _, board := range boards {
			board.Mark(number)
			if board.IsBingo() {
				return board.SumOfUnmarked() * number
			}
		}
	}
	return 0
}

func Part2(sequence []int, boards []types.Bingo) int {
	noBoards := len(boards)
	completedBoards := 0
	for _, number := range sequence {
		for index, board := range boards {
			if board.Completed {
				continue
			}
			board.Mark(number)
			if board.IsBingo() {
				boards[index].Completed = true
				completedBoards++
			}
			if completedBoards == noBoards {
				return board.SumOfUnmarked() * number
			}
		}
	}
	return 0
}
