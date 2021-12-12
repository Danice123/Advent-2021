package day11

import (
	"advent2021/day05"
	"advent2021/helper"
	"fmt"
)

func printmap(octomap map[int]map[int]int, flashes []day05.Coord) {
	colorRed := "\033[31m"
	colorReset := "\033[0m"
	for y := 0; y < len(octomap); y++ {
	xloop:
		for x := 0; x < len(octomap[0]); x++ {
			for _, c := range flashes {
				if c.X == x && c.Y == y {
					fmt.Printf("%s%d%s", colorRed, octomap[y][x], colorReset)
					continue xloop
				}
			}
			fmt.Printf("%d", octomap[y][x])
		}
		print("\n")
	}
}

func Part1() {
	input := helper.ReadInput("day11/input.txt")

	octomap := make(map[int]map[int]int)
	for y, line := range input {
		octomap[y] = make(map[int]int)
		for x, v := range line {
			octomap[y][x] = int(v - '0')
		}
	}

	flashTimes := 0
	for step := 0; step < 100; step++ {
		flashes := []day05.Coord{}
		for y := 0; y < len(octomap); y++ {
			for x := 0; x < len(octomap[0]); x++ {
				octomap[y][x]++
				if octomap[y][x] > 9 {
					flashes = append(flashes, day05.Coord{
						X: x,
						Y: y,
					})
					octomap[y][x] = 0
					flashTimes++
				}
			}
		}

		for len(flashes) > 0 {
			newFlashes := []day05.Coord{}

			checkflash := func(x int, y int) {
				if x < 0 || x >= len(octomap[0]) {
					return
				}
				if y < 0 || y >= len(octomap) {
					return
				}
				if octomap[y][x] == 0 {
					return
				}
				octomap[y][x]++
				if octomap[y][x] > 9 {
					newFlashes = append(newFlashes, day05.Coord{
						X: x,
						Y: y,
					})
					octomap[y][x] = 0
					flashTimes++
				}
			}

			for _, c := range flashes {
				checkflash(c.X+1, c.Y)
				checkflash(c.X-1, c.Y)
				checkflash(c.X, c.Y+1)
				checkflash(c.X, c.Y-1)
				checkflash(c.X+1, c.Y+1)
				checkflash(c.X+1, c.Y-1)
				checkflash(c.X-1, c.Y+1)
				checkflash(c.X-1, c.Y-1)
			}

			flashes = newFlashes
		}
	}

	println(flashTimes)
}

func Part2() {
	input := helper.ReadInput("day11/input.txt")

	octomap := make(map[int]map[int]int)
	for y, line := range input {
		octomap[y] = make(map[int]int)
		for x, v := range line {
			octomap[y][x] = int(v - '0')
		}
	}

	allzero := false
	for step := 0; !allzero; step++ {
		flashes := []day05.Coord{}
		for y := 0; y < len(octomap); y++ {
			for x := 0; x < len(octomap[0]); x++ {
				octomap[y][x]++
				if octomap[y][x] > 9 {
					flashes = append(flashes, day05.Coord{
						X: x,
						Y: y,
					})
					octomap[y][x] = 0
				}
			}
		}

		for len(flashes) > 0 {
			newFlashes := []day05.Coord{}

			checkflash := func(x int, y int) {
				if x < 0 || x >= len(octomap[0]) {
					return
				}
				if y < 0 || y >= len(octomap) {
					return
				}
				if octomap[y][x] == 0 {
					return
				}
				octomap[y][x]++
				if octomap[y][x] > 9 {
					newFlashes = append(newFlashes, day05.Coord{
						X: x,
						Y: y,
					})
					octomap[y][x] = 0
				}
			}

			for _, c := range flashes {
				checkflash(c.X+1, c.Y)
				checkflash(c.X-1, c.Y)
				checkflash(c.X, c.Y+1)
				checkflash(c.X, c.Y-1)
				checkflash(c.X+1, c.Y+1)
				checkflash(c.X+1, c.Y-1)
				checkflash(c.X-1, c.Y+1)
				checkflash(c.X-1, c.Y-1)
			}

			flashes = newFlashes
		}

		checkzero := true
	loop:
		for y := 0; y < len(octomap); y++ {
			for x := 0; x < len(octomap[0]); x++ {
				if octomap[y][x] != 0 {
					checkzero = false
					break loop
				}
			}
		}
		if checkzero {
			allzero = true
			println(step + 1)
		}
	}
}
