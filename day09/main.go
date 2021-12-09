package day09

import (
	"advent2021/day05"
	"advent2021/helper"
	"sort"
)

type HeightMap struct {
	Heights [][]int
}

func NewHeightMap(input []string) HeightMap {
	m := HeightMap{
		Heights: make([][]int, len(input)),
	}

	for i, line := range input {
		m.Heights[i] = make([]int, len(line))

		for j, c := range line {
			m.Heights[i][j] = int(c - '0')
		}
	}

	return m
}

func (ths HeightMap) FindLowPoints() []day05.Coord {
	lowPoints := []day05.Coord{}
	for i, row := range ths.Heights {
		for j, v := range row {
			isLowPoint := true
			if i != 0 && v >= ths.Heights[i-1][j] {
				isLowPoint = false
			}
			if i != len(ths.Heights)-1 && v >= ths.Heights[i+1][j] {
				isLowPoint = false
			}
			if j != 0 && v >= ths.Heights[i][j-1] {
				isLowPoint = false
			}
			if j != len(row)-1 && v >= ths.Heights[i][j+1] {
				isLowPoint = false
			}
			if isLowPoint {
				lowPoints = append(lowPoints, day05.Coord{
					X: j,
					Y: i,
				})
			}
		}
	}
	return lowPoints
}

func (ths HeightMap) CalculateBasin(lowPoint day05.Coord) int {
	basin := map[day05.Coord]bool{
		lowPoint: false,
	}

	for {
		newPointsAdded := false
		for c, hasBeenEvaluated := range basin {
			if hasBeenEvaluated {
				continue
			}
			if c.Y != 0 && ths.Heights[c.Y-1][c.X] < 9 {
				if _, ok :=
					basin[day05.Coord{
						X: c.X,
						Y: c.Y - 1,
					}]; !ok {
					basin[day05.Coord{
						X: c.X,
						Y: c.Y - 1,
					}] = false
				}
				newPointsAdded = true
			}
			if c.Y != len(ths.Heights)-1 && ths.Heights[c.Y+1][c.X] < 9 {
				if _, ok :=
					basin[day05.Coord{
						X: c.X,
						Y: c.Y + 1,
					}]; !ok {
					basin[day05.Coord{
						X: c.X,
						Y: c.Y + 1,
					}] = false
				}
				newPointsAdded = true
			}
			if c.X != 0 && ths.Heights[c.Y][c.X-1] < 9 {
				if _, ok :=
					basin[day05.Coord{
						X: c.X - 1,
						Y: c.Y,
					}]; !ok {
					basin[day05.Coord{
						X: c.X - 1,
						Y: c.Y,
					}] = false
				}
				newPointsAdded = true
			}
			if c.X != len(ths.Heights[c.Y])-1 && ths.Heights[c.Y][c.X+1] < 9 {
				if _, ok :=
					basin[day05.Coord{
						X: c.X + 1,
						Y: c.Y,
					}]; !ok {
					basin[day05.Coord{
						X: c.X + 1,
						Y: c.Y,
					}] = false
				}
				newPointsAdded = true
			}
			basin[c] = true
		}
		if !newPointsAdded {
			break
		}
	}

	return len(basin)
}

func Part1() {
	input := helper.ReadInput("day09/input.txt")

	hm := NewHeightMap(input)
	lp := hm.FindLowPoints()

	sumOfDanger := 0
	for _, p := range lp {
		sumOfDanger += hm.Heights[p.Y][p.X] + 1
	}

	println(sumOfDanger)
}

func Part2() {
	input := helper.ReadInput("day09/input.txt")
	hm := NewHeightMap(input)
	lp := hm.FindLowPoints()

	sizes := []int{}
	for _, p := range lp {
		sizes = append(sizes, hm.CalculateBasin(p))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	println(sizes[0] * sizes[1] * sizes[2])
}
