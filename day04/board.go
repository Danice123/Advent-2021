package day04

import (
	"strconv"
	"strings"
)

type BingoBoard struct {
	numbers [][]int
	marked  [][]bool
}

func (ths *BingoBoard) MarkNumber(number int) bool {
	for i, line := range ths.numbers {
		for j, n := range line {
			if n == number {
				ths.marked[i][j] = true
			}
		}
	}
	return ths.check()
}

func (ths *BingoBoard) check() bool {
	for i := 0; i < 5; i++ {
		across := true
		down := true
		diag := true
		for j := 0; j < 5; j++ {
			// check rows
			if !ths.marked[i][j] {
				across = false
			}
			// check columns
			if !ths.marked[j][i] {
				down = false
			}
			// check diagonals
			if i == 0 && !ths.marked[j][j] {
				diag = false
			} else if i == 1 && !ths.marked[j][4-j] {
				diag = false
			} else {
				diag = false
			}
		}
		if across || down || diag {
			return true
		}
	}
	return false
}

func (ths *BingoBoard) Sum() int {
	sum := 0
	for i, line := range ths.marked {
		for j, m := range line {
			if !m {
				sum += ths.numbers[i][j]
			}
		}
	}
	return sum
}

func NewBingoBoard(data []byte) BingoBoard {
	board := BingoBoard{
		numbers: make([][]int, 5),
		marked:  make([][]bool, 5),
	}

	for i, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}
		if i > 4 {
			panic("Board input is bad")
		}

		for _, numberstring := range strings.Split(line, " ") {
			if numberstring == "" {
				continue
			}
			number, err := strconv.Atoi(numberstring)
			if err != nil {
				panic(err)
			}
			board.numbers[i] = append(board.numbers[i], number)
			board.marked[i] = append(board.marked[i], false)
		}

		if len(board.numbers[i]) != 5 || len(board.marked[i]) != 5 {
			panic("Board is bad")
		}
	}

	return board
}
