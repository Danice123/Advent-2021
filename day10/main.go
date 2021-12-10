package day10

import (
	"advent2021/helper"
	"fmt"
	"sort"
)

func Part1() {
	lines := helper.ReadInput("day10/input.txt")

	points := 0
	for _, line := range lines {
		stack := []rune{}
	reader:
		for i, r := range line {
			switch r {
			case '(':
				fallthrough
			case '[':
				fallthrough
			case '{':
				fallthrough
			case '<':
				stack = append(stack, r)
			case ')':
				if len(stack) == 0 {
					continue
				}
				index := len(stack) - 1
				previous := stack[index]
				stack = stack[:index]
				if previous != '(' {
					points += 3
					fmt.Printf("Corrupted line: %s, col: %d, expected %c but found %c\n", line, i, previous, ')')
					break reader
				}
			case ']':
				if len(stack) == 0 {
					continue
				}
				index := len(stack) - 1
				previous := stack[index]
				stack = stack[:index]
				if previous != '[' {
					points += 57
					fmt.Printf("Corrupted line: %s, col: %d, expected %c but found %c\n", line, i, previous, ']')
					break reader
				}
			case '}':
				if len(stack) == 0 {
					continue
				}
				index := len(stack) - 1
				previous := stack[index]
				stack = stack[:index]
				if previous != '{' {
					points += 1197
					fmt.Printf("Corrupted line: %s, col: %d, expected %c but found %c\n", line, i, previous, '}')
					break reader
				}
			case '>':
				if len(stack) == 0 {
					continue
				}
				index := len(stack) - 1
				previous := stack[index]
				stack = stack[:index]
				if previous != '<' {
					points += 25137
					fmt.Printf("Corrupted line: %s, col: %d, expected %c but found %c\n", line, i, previous, '>')
					break reader
				}
			default:
				fmt.Printf("Wierd char: %c", r)
			}
		}
	}

	println(points)
}

func Part2() {
	lines := helper.ReadInput("day10/input.txt")

	scoreList := []int{}
	for _, line := range lines {
		points := 0

		stack := []rune{}
		corrupt := false
	reader:
		for _, r := range line {
			switch r {
			case '(':
				fallthrough
			case '[':
				fallthrough
			case '{':
				fallthrough
			case '<':
				stack = append(stack, r)
			case ')':
				if len(stack) == 0 {
					continue
				}
				index := len(stack) - 1
				previous := stack[index]
				stack = stack[:index]
				if previous != '(' {
					corrupt = true
					break reader
				}
			case ']':
				if len(stack) == 0 {
					continue
				}
				index := len(stack) - 1
				previous := stack[index]
				stack = stack[:index]
				if previous != '[' {
					corrupt = true
					break reader
				}
			case '}':
				if len(stack) == 0 {
					continue
				}
				index := len(stack) - 1
				previous := stack[index]
				stack = stack[:index]
				if previous != '{' {
					corrupt = true
					break reader
				}
			case '>':
				if len(stack) == 0 {
					continue
				}
				index := len(stack) - 1
				previous := stack[index]
				stack = stack[:index]
				if previous != '<' {
					corrupt = true
					break reader
				}
			}
		}

		if !corrupt {
			fix := ""
			for i := len(stack) - 1; i >= 0; i-- {
				switch stack[i] {
				case '(':
					fix += ")"
					points = points*5 + 1
				case '[':
					fix += "]"
					points = points*5 + 2
				case '{':
					fix += "}"
					points = points*5 + 3
				case '<':
					fix += ">"
					points = points*5 + 4
				}
			}
			fmt.Printf("Fix %s with %s\n", line, fix)
			println(points)
			scoreList = append(scoreList, points)
		}
	}

	sort.Ints(scoreList)
	println(scoreList[len(scoreList)/2])
}
