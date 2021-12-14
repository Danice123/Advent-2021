package day14

import (
	"advent2021/helper"
	"sort"
	"strings"
)

func Part1() {
	input := helper.ReadInput("day14/input.txt")

	var polymer string
	rules := map[string]string{}
	for i, line := range input {
		if line == "" {
			continue
		}
		if i == 0 {
			polymer = line
			continue
		}

		rule := strings.Split(line, " -> ")
		rules[rule[0]] = rule[1]
	}

	for i := 0; i < 10; i++ {
		for w := len(polymer) - 2; w >= 0; w-- {
			if insert, ok := rules[string([]byte{polymer[w], polymer[w+1]})]; ok {
				// fmt.Printf("%s => %s + %s + %s\n", pair, polymer[:w+1], insert, polymer[w+1:])
				polymer = polymer[:w+1] + insert + polymer[w+1:]
			}
		}
	}

	elements := map[rune]int{}
	for _, r := range polymer {
		if _, ok := elements[r]; !ok {
			elements[r] = 1
		} else {
			elements[r]++
		}
	}

	values := []int{}
	for _, v := range elements {
		values = append(values, v)
	}
	sort.Ints(values)

	println(values[len(values)-1] - values[0])
}

func Part2() {
	input := helper.ReadInput("day14/input.txt")

	polymer := map[string]int{}
	polymerRules := map[string][]string{}
	counts := map[rune]int{}
	countRules := map[string]rune{}
	for i, line := range input {
		if line == "" {
			continue
		}
		if i == 0 {
			for w := 0; w < len(line)-1; w++ {
				polymer[string([]byte{line[w], line[w+1]})]++
			}
			for _, r := range line {
				counts[r]++
			}
			continue
		}

		rule := strings.Split(line, " -> ")
		polymerRules[rule[0]] = []string{
			string([]byte{rule[0][0], rule[1][0]}),
			string([]byte{rule[1][0], rule[0][1]}),
		}
		countRules[rule[0]] = rune(rule[1][0])
	}

	for i := 0; i < 40; i++ {
		newPolymer := map[string]int{}
		for pair, n := range polymer {
			if add, ok := polymerRules[pair]; ok {
				for _, p := range add {
					newPolymer[p] += n
				}
				counts[countRules[pair]] += n
			}
		}
		polymer = newPolymer
	}

	values := []int{}
	for _, n := range counts {
		values = append(values, n)
	}
	sort.Ints(values)
	println(values[len(values)-1] - values[0])
}
