package types

import "fmt"

type BoardCell struct {
	Value  int
	Marked bool
}

type Bingo struct {
	Board     [][]BoardCell
	Completed bool
}

func (board *Bingo) Mark(marker int) {
	for i := 0; i < len(board.Board); i++ {
		for j := 0; j < len(board.Board[i]); j++ {
			if board.Board[i][j].Value == marker {
				board.Board[i][j].Marked = true
			}
		}
	}
}

func (board *Bingo) IsBingo() bool {
	for i := 0; i < len(board.Board); i++ {
		if board.isRowMarked(i) {
			board.Completed = true
			return true
		}
	}
	for i := 0; i < len(board.Board[0]); i++ {
		if board.isColumnMarked(i) {
			board.Completed = true
			return true
		}
	}
	return false
}

func (board *Bingo) SumOfUnmarked() int {
	sum := 0
	for i := 0; i < len(board.Board); i++ {
		for j := 0; j < len(board.Board[i]); j++ {
			if !board.Board[i][j].Marked {
				sum += board.Board[i][j].Value
			}
		}
	}
	return sum
}

func (board *Bingo) isRowMarked(row int) bool {
	for i := 0; i < len(board.Board[row]); i++ {
		if !board.Board[row][i].Marked {
			return false
		}
	}
	return true
}

func (board *Bingo) isColumnMarked(column int) bool {
	for i := 0; i < len(board.Board); i++ {
		if !board.Board[i][column].Marked {
			return false
		}
	}
	return true
}

func (board *Bingo) Print() {
	for i := 0; i < len(board.Board); i++ {
		for j := 0; j < len(board.Board[i]); j++ {
			if board.Board[i][j].Marked {
				fmt.Print(" X")
			} else {
				fmt.Print(" ", board.Board[i][j].Value)
			}
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}
