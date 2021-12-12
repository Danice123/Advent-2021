package day12

import (
	"advent2021/helper"
	"strings"
)

type Cave struct {
	Name        string
	Small       bool
	Connections []*Cave
}

func NewCave(name string) *Cave {
	return &Cave{
		Name:        name,
		Small:       strings.ToUpper(name) != name,
		Connections: []*Cave{},
	}
}

func BuildCaves(input []string) map[string]*Cave {
	caves := map[string]*Cave{}
	for _, line := range input {
		s := strings.Split(line, "-")

		if _, ok := caves[s[0]]; !ok {
			caves[s[0]] = NewCave(s[0])
		}
		if _, ok := caves[s[1]]; !ok {
			caves[s[1]] = NewCave(s[1])
		}
		caves[s[0]].Connections = append(caves[s[0]].Connections, caves[s[1]])
		caves[s[1]].Connections = append(caves[s[1]].Connections, caves[s[0]])
	}
	return caves
}

type Path []string

func (ths Path) Current() string {
	return ths[len(ths)-1]
}

func (ths Path) PotentialNext(current *Cave) []string {
	visited := make(map[string]bool)
	for _, cave := range ths {
		visited[cave] = true
	}

	potentials := []string{}
	for _, cave := range current.Connections {
		if cave.Name == "start" {
			continue
		}
		if cave.Small && visited[cave.Name] {
			continue
		}
		potentials = append(potentials, cave.Name)
	}
	return potentials
}

func (ths Path) PotentialNextTwo(current *Cave) []string {
	visited := make(map[string]int)
	hasVisitedSmallCaveTwice := false
	for _, cave := range ths {
		if _, ok := visited[cave]; !ok {
			visited[cave] = 1
		} else {
			if strings.ToUpper(cave) != cave {
				hasVisitedSmallCaveTwice = true
			}
			visited[cave]++
		}
	}

	potentials := []string{}
	for _, cave := range current.Connections {
		if cave.Name == "start" {
			continue
		}
		if cave.Small {
			if hasVisitedSmallCaveTwice && visited[cave.Name] > 0 {
				continue
			}
			if !hasVisitedSmallCaveTwice && visited[cave.Name] > 1 {
				continue
			}
		}
		potentials = append(potentials, cave.Name)
	}
	return potentials
}

func Part1() {
	input := helper.ReadInput("day12/input.txt")
	caves := BuildCaves(input)

	paths := []Path{
		{"start"},
	}
	terminated := []Path{}

	for len(paths) > 0 {
		newPaths := []Path{}
		for _, path := range paths {
			potentials := path.PotentialNext(caves[path.Current()])
			for _, potential := range potentials {
				if potential == "end" {
					np := append(Path{}, path...)
					terminated = append(terminated, append(np, potential))
				} else {
					np := append(Path{}, path...)
					newPaths = append(newPaths, append(np, potential))
				}
			}
		}
		paths = newPaths
	}
	println(len(terminated))
}

func Part2() {
	input := helper.ReadInput("day12/input.txt")
	caves := BuildCaves(input)

	paths := []Path{
		{"start"},
	}
	terminated := []Path{}

	for len(paths) > 0 {
		newPaths := []Path{}
		for _, path := range paths {
			potentials := path.PotentialNextTwo(caves[path.Current()])
			for _, potential := range potentials {
				if potential == "end" {
					np := append(Path{}, path...)
					terminated = append(terminated, append(np, potential))
				} else {
					np := append(Path{}, path...)
					newPaths = append(newPaths, append(np, potential))
				}
			}
		}
		paths = newPaths
	}

	println(len(terminated))
}
