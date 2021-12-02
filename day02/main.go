package day02

import (
	"advent2021/helper"
	"regexp"
	"strconv"
)

var regex = regexp.MustCompile(`^(\D*) (\d*)$`)

func Part1() {
	instructions := helper.ReadInput("day02/input.txt")

	horz := 0
	depth := 0
	for _, instruct := range instructions {
		match := regex.FindAllStringSubmatch(instruct, -1)[0]
		value, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		switch match[1] {
		case "forward":
			horz += value
		case "down":
			depth += value
		case "up":
			depth -= value
		default:
			panic("Bad command " + match[0])
		}
	}

	println(horz * depth)
}

func Part2() {
	instructions := helper.ReadInput("day02/input.txt")

	horz := 0
	depth := 0
	aim := 0
	for _, instruct := range instructions {
		match := regex.FindAllStringSubmatch(instruct, -1)[0]
		value, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		switch match[1] {
		case "forward":
			horz += value
			depth += value * aim
		case "down":
			aim += value
		case "up":
			aim -= value
		default:
			panic("Bad command " + match[0])
		}
	}

	println(horz * depth)
}
