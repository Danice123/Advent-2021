package day15

import (
	"advent2021/helper"
	"fmt"
)

type Tile struct {
	X int
	Y int
}

func (ths Tile) GetNeighbors(maxX int, maxY int) []Tile {
	neighbors := []Tile{}
	if ths.X > 0 {
		neighbors = append(neighbors, Tile{X: ths.X - 1, Y: ths.Y})
	}
	if ths.X < maxX {
		neighbors = append(neighbors, Tile{X: ths.X + 1, Y: ths.Y})
	}
	if ths.Y < maxY {
		neighbors = append(neighbors, Tile{X: ths.X, Y: ths.Y + 1})
	}
	if ths.Y > 0 {
		neighbors = append(neighbors, Tile{X: ths.X, Y: ths.Y - 1})
	}
	return neighbors
}

type Cave struct {
	cost map[int]map[int]int
}

func NewCave(input []string) *Cave {
	cave := &Cave{
		cost: map[int]map[int]int{},
	}
	for y, line := range input {
		cave.cost[y] = map[int]int{}
		for x, r := range line {
			cave.cost[y][x] = int(r - '0')
		}
	}
	return cave
}

func NewTiledCave(tile *Cave) *Cave {
	cave := &Cave{
		cost: map[int]map[int]int{},
	}

	for tileY := 0; tileY < 5; tileY++ {
		for tileX := 0; tileX < 5; tileX++ {
			additionalCost := tileY + tileX
			for y := 0; y < len(tile.cost); y++ {
				if _, ok := cave.cost[tileY*len(tile.cost)+y]; !ok {
					cave.cost[tileY*len(tile.cost)+y] = map[int]int{}
				}
				for x := 0; x < len(tile.cost[0]); x++ {
					newCost := tile.cost[y][x] + additionalCost
					if newCost > 9 {
						newCost = newCost%10 + 1
					}
					cave.cost[tileY*len(tile.cost)+y][tileX*len(tile.cost[0])+x] = newCost
				}
			}
		}
	}
	return cave
}

// https://en.wikipedia.org/wiki/A*_search_algorithm
func (ths *Cave) AStar() []Tile {
	start := Tile{X: 0, Y: 0}
	goal := Tile{X: len(ths.cost[0]) - 1, Y: len(ths.cost) - 1}

	// The set of discovered nodes that may need to be (re-)expanded.
	// Initially, only the start node is known.
	// This is usually implemented as a min-heap or priority queue rather than a hash-set.
	openSet := map[Tile]bool{start: true}

	// For node n, cameFrom[n] is the node immediately preceding it on the cheapest path from start
	// to n currently known.
	cameFrom := map[Tile]Tile{}

	// For node n, gScore[n] is the cost of the cheapest path from start to n currently known.
	gScore := map[Tile]int{} // with default value of Infinity
	gScore[start] = 0

	// For node n, fScore[n] := gScore[n] + h(n). fScore[n] represents our current best guess as to
	// how short a path from start to finish can be if it goes through n.
	fScore := map[Tile]int{} // with default value of Infinity
	fScore[start] = 0

	for len(openSet) > 0 {
		// This operation can occur in O(1) time if openSet is a min-heap or a priority queue
		var current Tile // current := the node in openSet having the lowest fScore[] value
		currentValue := -1
		for c := range openSet {
			cValue := fScore[c]
			if currentValue == -1 || currentValue > cValue {
				current = c
				currentValue = cValue
			}
		}

		if current == goal {
			return reconstruct_path(cameFrom, current)
		}

		delete(openSet, current)
		for _, neighbor := range current.GetNeighbors(len(ths.cost[0])-1, len(ths.cost)-1) {
			// d(current,neighbor) is the weight of the edge from current to neighbor
			// tentative_gScore is the distance from start to the neighbor through current
			tentative_gScore := gScore[current] + ths.cost[neighbor.Y][neighbor.X]
			if v, ok := gScore[neighbor]; tentative_gScore < v || !ok {
				// This path to neighbor is better than any previous one. Record it!
				cameFrom[neighbor] = current
				gScore[neighbor] = tentative_gScore
				fScore[neighbor] = tentative_gScore // + ths.cost[neighbor.Y][neighbor.X]
				if _, ok := openSet[neighbor]; !ok {
					openSet[neighbor] = true
				}
			}
		}
	}

	// Open set is empty but goal was never reached
	return nil
}

func reconstruct_path(cameFrom map[Tile]Tile, current Tile) []Tile {
	total_path := []Tile{current}
	for {
		if _, ok := cameFrom[current]; !ok {
			break
		}
		current = cameFrom[current]
		total_path = append([]Tile{current}, total_path...)
	}
	return total_path
}

func (ths *Cave) Print(path []Tile) {
	colorRed := "\033[31m"
	colorReset := "\033[0m"

	pathMap := map[Tile]bool{}
	for _, t := range path {
		pathMap[t] = true
	}
	for y := 0; y < len(ths.cost); y++ {
		for x := 0; x < len(ths.cost[0]); x++ {
			if _, ok := pathMap[Tile{X: x, Y: y}]; ok {
				fmt.Print(string(colorRed))
				fmt.Printf("%d", ths.cost[y][x])
				fmt.Print(string(colorReset))
			} else {
				fmt.Printf("%d", ths.cost[y][x])
			}
		}
		fmt.Print("\n")
	}
}

func Part1() {
	input := helper.ReadInput("day15/input.txt")

	cave := NewCave(input)
	path := cave.AStar()
	fmt.Printf("%v\n", path)

	cave.Print(path)

	sum := 0
	for i, p := range path {
		if i == 0 {
			continue
		}
		sum += cave.cost[p.Y][p.X]
	}
	println(sum)
}

func Part2() {
	input := helper.ReadInput("day15/input.txt")

	tile := NewCave(input)
	cave := NewTiledCave(tile)
	path := cave.AStar()
	//cave.Print(path)

	sum := 0
	for i, p := range path {
		if i == 0 {
			continue
		}
		sum += cave.cost[p.Y][p.X]
	}
	println(sum)
}
