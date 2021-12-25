package day25

import (
	"advent2021/helper"
	"fmt"
)

type Coord struct {
	X int
	Y int
}

func Part1() {
	input := helper.ReadInput("day25/input.txt")

	var maxX int
	maxY := len(input)
	eastCucumbers := map[Coord]bool{}
	southCucumbers := map[Coord]bool{}
	for y, line := range input {
		maxX = len(line)
		for x, r := range line {
			if r == '>' {
				eastCucumbers[Coord{X: x, Y: y}] = true
			}
			if r == 'v' {
				southCucumbers[Coord{X: x, Y: y}] = true
			}
		}
	}

	step := 0
	for {
		var moved bool

		newEast := map[Coord]bool{}
		for cuce := range eastCucumbers {
			move := Coord{X: cuce.X + 1, Y: cuce.Y}
			if move.X == maxX {
				move = Coord{X: 0, Y: cuce.Y}
			}
			if !eastCucumbers[move] && !southCucumbers[move] {
				newEast[move] = true
				moved = true
			} else {
				newEast[cuce] = true
			}
		}
		eastCucumbers = newEast

		newSouth := map[Coord]bool{}
		for cuce := range southCucumbers {
			move := Coord{X: cuce.X, Y: cuce.Y + 1}
			if move.Y == maxY {
				move = Coord{X: cuce.X, Y: 0}
			}
			if !eastCucumbers[move] && !southCucumbers[move] {
				newSouth[move] = true
				moved = true
			} else {
				newSouth[cuce] = true
			}
		}
		southCucumbers = newSouth

		step++
		if !moved {
			fmt.Printf("Step: %d\n", step)
			for y := 0; y < maxY; y++ {
				for x := 0; x < maxX; x++ {
					c := Coord{X: x, Y: y}
					if eastCucumbers[c] {
						print(">")
					} else if southCucumbers[c] {
						print("v")
					} else {
						print(".")
					}
				}
				println("")
			}
			break
		}
	}
}
