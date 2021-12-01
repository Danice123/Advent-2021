package day1

import (
	"os"
	"strconv"
	"strings"
)

func readInput() []string {
	data, err := os.ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

func Part1() {
	depths := readInput()
	depthList := make([]int, len(depths))
	increases := 0
	for i, depth := range depths {
		if d, err := strconv.Atoi(depth); err != nil {
			panic(err)
		} else {
			depthList[i] = d
		}
		if i > 0 && depthList[i-1] < depthList[i] {
			increases++
		}
	}

	println(increases)
}

func Part2() {
	depths := readInput()
	depthList := make([]int, len(depths))
	increases := 0
	for i, depth := range depths {
		if d, err := strconv.Atoi(depth); err != nil {
			panic(err)
		} else {
			depthList[i] = d
		}

		if i > 2 {
			previousSum := depthList[i-3] + depthList[i-2] + depthList[i-1]
			sum := depthList[i] + depthList[i-1] + depthList[i-2]
			if sum > previousSum {
				increases++
			}
		}
	}

	println(increases)
}
