package day08

import (
	"advent2021/helper"
	"sort"
	"strconv"
	"strings"
)

func Part1() {
	lines := helper.ReadInput("day08/input.txt")

	count := 0
	for _, line := range lines {
		outputs := strings.TrimSpace(strings.Split(line, "|")[1])

		for _, output := range strings.Split(outputs, " ") {
			if len(output) == 2 || len(output) == 3 || len(output) == 4 || len(output) == 7 {
				count++
			}
		}
	}
	println(count)
}

func Part2() {
	lines := helper.ReadInput("day08/input.txt")

	total := 0
	for _, line := range lines {
		cypher := NewCypher()

		input := strings.Split(strings.TrimSpace(strings.Split(line, "|")[0]), " ")
		sort.Slice(input, func(i int, j int) bool {
			return len(input[i]) < len(input[j])
		})

		for _, inputWord := range input {
			possible := make(map[rune]int)
			impossible := map[rune]int{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0}
			for _, c := range inputWord {
				possible[c] = 0
				delete(impossible, c)
			}

			switch len(inputWord) {
			case 3: // digit 7
				cypher.NarrowPossibleOptions('a', possible)
				cypher.NarrowPossibleOptions('b', impossible)
				cypher.NarrowPossibleOptions('c', possible)
				cypher.NarrowPossibleOptions('d', impossible)
				cypher.NarrowPossibleOptions('e', impossible)
				cypher.NarrowPossibleOptions('f', possible)
				cypher.NarrowPossibleOptions('g', impossible)
			case 2: // digit 1
				cypher.NarrowPossibleOptions('a', impossible)
				cypher.NarrowPossibleOptions('b', impossible)
				cypher.NarrowPossibleOptions('c', possible)
				cypher.NarrowPossibleOptions('d', impossible)
				cypher.NarrowPossibleOptions('e', impossible)
				cypher.NarrowPossibleOptions('f', possible)
				cypher.NarrowPossibleOptions('g', impossible)
			case 4: // digit 4
				cypher.NarrowPossibleOptions('a', impossible)
				cypher.NarrowPossibleOptions('b', possible)
				cypher.NarrowPossibleOptions('c', possible)
				cypher.NarrowPossibleOptions('d', possible)
				cypher.NarrowPossibleOptions('e', impossible)
				cypher.NarrowPossibleOptions('f', possible)
				cypher.NarrowPossibleOptions('g', impossible)
			case 5: // digits 2, 3, 5

			case 6: // digits 0, 6, 9
				var leftOver rune
				for c := range impossible {
					leftOver = c
					break
				}

				if cypher.Has('c', leftOver) && !cypher.Has('d', leftOver) && !cypher.Has('e', leftOver) {
					cypher.NarrowPossibleOptions('c', map[rune]int{leftOver: 0})
					break
				}
				if !cypher.Has('c', leftOver) && cypher.Has('d', leftOver) && !cypher.Has('e', leftOver) {
					cypher.NarrowPossibleOptions('d', map[rune]int{leftOver: 0})
					break
				}
				if !cypher.Has('c', leftOver) && !cypher.Has('d', leftOver) && cypher.Has('e', leftOver) {
					cypher.NarrowPossibleOptions('e', map[rune]int{leftOver: 0})
				}
			}
			if cypher.IsSolved() {
				break
			}
		}

		output := strings.Split(strings.TrimSpace(strings.Split(line, "|")[1]), " ")
		numstr := cypher.Solve(output)
		num, err := strconv.Atoi(numstr)
		if err != nil {
			panic(err)
		}
		total += num
	}

	println(total)
}
