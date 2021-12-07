package day07

import (
	"advent2021/helper"
	"fmt"
	"math"
)

func Part1() {
	crabs := helper.ReadNumberList("day07/test.txt")
	crabmap := make(map[int]int)
	max := 1
	for _, p := range crabs {
		if p > max {
			max = p
		}
		if _, ok := crabmap[p]; !ok {
			crabmap[p] = 0
		}
		crabmap[p]++
	}

	bestCost := -1
	bestPos := -1
	for i := 0; i < max; i++ {
		cost := 0
		for p, crabs := range crabmap {
			distance := int(math.Abs(float64(i - p)))
			cost += distance * crabs
		}
		if cost < bestCost || bestCost == -1 {
			bestCost = cost
			bestPos = i
		}
	}

	fmt.Printf("Position %d: cost %d", bestPos, bestCost)
}

func Part2() {
	crabs := helper.ReadNumberList("day07/input.txt")
	crabmap := make(map[int]int)
	max := 1
	for _, p := range crabs {
		if p > max {
			max = p
		}
		if _, ok := crabmap[p]; !ok {
			crabmap[p] = 0
		}
		crabmap[p]++
	}

	bestCost := -1
	bestPos := -1
	for i := 0; i < max; i++ {
		cost := 0
		for p, crabs := range crabmap {
			distance := int(math.Abs(float64(i - p)))
			if distance%2 == 1 { //odd
				cost += int(math.Ceil(float64(distance)/2.0)) * distance * crabs
			} else {
				cost += (distance*distance/2 + distance/2) * crabs
			}
		}
		if cost < bestCost || bestCost == -1 {
			bestCost = cost
			bestPos = i
		}
	}

	fmt.Printf("Position %d: cost %d", bestPos, bestCost)
}
