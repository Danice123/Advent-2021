package day05

import (
	"os"
	"regexp"
)

var lineRegex = regexp.MustCompile(`(\d+,\d+) -> (\d+,\d+)`)

func Part1() {
	data, err := os.ReadFile("day05/input.txt")
	if err != nil {
		panic(err)
	}
	matches := lineRegex.FindAllSubmatch(data, -1)
	if matches == nil {
		panic("Bad input")
	}

	lines := []Line{}
	for _, input := range matches {
		line := NewLine(input)
		if line.Start.X == line.End.X || line.Start.Y == line.End.Y {
			lines = append(lines, line)
		}
	}

	lineMap := make(map[Coord]int)
	for _, line := range lines {
		for _, c := range line.GetCoverage() {
			if _, ok := lineMap[c]; !ok {
				lineMap[c] = 1
			} else {
				lineMap[c]++
			}
		}
	}

	overlaps := 0
	for _, o := range lineMap {
		if o > 1 {
			overlaps++
		}
	}

	println(overlaps)
}

func Part2() {
	data, err := os.ReadFile("day05/input.txt")
	if err != nil {
		panic(err)
	}
	matches := lineRegex.FindAllSubmatch(data, -1)
	if matches == nil {
		panic("Bad input")
	}

	lines := []Line{}
	for _, input := range matches {
		lines = append(lines, NewLine(input))
	}

	lineMap := make(map[Coord]int)
	for _, line := range lines {
		for _, c := range line.GetCoverage() {
			if _, ok := lineMap[c]; !ok {
				lineMap[c] = 1
			} else {
				lineMap[c]++
			}
		}
	}

	overlaps := 0
	for _, o := range lineMap {
		if o > 1 {
			overlaps++
		}
	}

	println(overlaps)
}
