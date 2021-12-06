package day04

import (
	"os"
	"regexp"
)

var bingoBoardRegex = regexp.MustCompile(`((?:[\d ]*\n){5})\n`)

func buildBoards(data []byte) []*BingoBoard {
	boardData := bingoBoardRegex.FindAllSubmatch(data, -1)

	boards := []*BingoBoard{}
	for _, boardData := range boardData {
		board := NewBingoBoard(boardData[1])
		boards = append(boards, &board)
	}

	return boards
}

func remove(s []*BingoBoard, i int) []*BingoBoard {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func Part1() {
	data, err := os.ReadFile("day04/input.txt")
	if err != nil {
		panic(err)
	}
	iterator := NewNumberIterator(data)
	boards := buildBoards(data)

	var winner *BingoBoard
	var winningNumber int
main:
	for {
		n := iterator.Next()

		for _, b := range boards {
			if b.MarkNumber(n) {
				winner = b
				winningNumber = n
				break main
			}
		}
	}

	println(winner.Sum() * winningNumber)
}

func Part2() {
	data, err := os.ReadFile("day04/input.txt")
	if err != nil {
		panic(err)
	}
	iterator := NewNumberIterator(data)
	boards := buildBoards(data)

	var winner *BingoBoard
	var winningNumber int
main:
	for {
		n := iterator.Next()

		removedBoards := []*BingoBoard{}
		for _, b := range boards {
			if b.MarkNumber(n) {
				if len(boards) == 1 {
					winner = b
					winningNumber = n
					break main
				}
				removedBoards = append(removedBoards, b)
			}
		}

		for _, toRemove := range removedBoards {
			for i, b := range boards {
				if b == toRemove {
					boards = remove(boards, i)
				}
			}

		}
	}

	println(winner.Sum() * winningNumber)
}
