package day13

import (
	"advent2021/day05"
	"advent2021/helper"
	"strconv"
	"strings"
)

type Axis string

const X_AXIS Axis = "x"
const Y_AXIS Axis = "y"

type Fold struct {
	axis  Axis
	value int
}

func ReadInput(input []string) ([]day05.Coord, []Fold) {
	dots := []day05.Coord{}
	folds := []Fold{}
	for _, line := range input {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "fold along") {
			c := strings.Split(strings.TrimPrefix(line, "fold along "), "=")
			v, err := strconv.Atoi(c[1])
			if err != nil {
				panic(err)
			}
			folds = append(folds, Fold{
				axis:  Axis(c[0]),
				value: v,
			})
		} else {
			c := strings.Split(line, ",")
			x, err := strconv.Atoi(c[0])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(c[1])
			if err != nil {
				panic(err)
			}
			dots = append(dots, day05.Coord{
				X: x,
				Y: y,
			})
		}
	}
	return dots, folds
}

func DoFold(dots []day05.Coord, fold Fold) []day05.Coord {
	newDots := []day05.Coord{}
	if fold.axis == X_AXIS {
		for _, dot := range dots {
			if dot.X < fold.value {
				newDots = append(newDots, dot)
			} else {
				newDots = append(newDots, day05.Coord{
					X: fold.value*2 - dot.X,
					Y: dot.Y,
				})
			}
		}
	} else if fold.axis == Y_AXIS {
		for _, dot := range dots {
			if dot.Y < fold.value {
				newDots = append(newDots, dot)
			} else {
				newDots = append(newDots, day05.Coord{
					X: dot.X,
					Y: fold.value*2 - dot.Y,
				})
			}
		}
	}
	return newDots
}

func Part1() {
	input := helper.ReadInput("day13/input.txt")
	dots, folds := ReadInput(input)
	dots = DoFold(dots, folds[0])

	overlap := make(map[day05.Coord]int)
	for _, dot := range dots {
		if _, ok := overlap[dot]; !ok {
			overlap[dot] = 1
		} else {
			overlap[dot]++
		}
	}

	println(len(overlap))

}

func Part2() {
	input := helper.ReadInput("day13/input.txt")
	dots, folds := ReadInput(input)

	for _, fold := range folds {
		dots = DoFold(dots, fold)
	}

	for y := 0; y < 8; y++ {
		for x := 0; x < 60; x++ {
			c := day05.Coord{
				X: x,
				Y: y,
			}
			found := false
			for _, dot := range dots {
				if dot == c {
					found = true
					break
				}
			}
			if found {
				print("#")
			} else {
				print(".")
			}
		}
		print("\n")
	}
}
