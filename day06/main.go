package day06

import (
	"advent2021/helper"
	"strconv"
	"strings"
)

func Part1() {
	data := helper.ReadInput("day06/input.txt")[0]

	fish := []int{}
	for _, s := range strings.Split(data, ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		fish = append(fish, i)
	}

	for t := 0; t < 80; t++ {
		newFish := []int{}
		for i, age := range fish {
			if age == 0 {
				newFish = append(newFish, 8)
				fish[i] = 6
			} else {
				fish[i]--
			}
		}
		fish = append(fish, newFish...)
	}

	println(len(fish))
}

func makeFishMap() map[int]int {
	return map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}
}

func Part2() {
	data := helper.ReadInput("day06/input.txt")[0]
	fish := makeFishMap()
	for _, s := range strings.Split(data, ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		fish[i]++
	}

	for t := 0; t < 256; t++ {
		newFish := makeFishMap()
		for age, n := range fish {
			if age == 0 {
				newFish[6] += n
				newFish[8] += n
			} else {
				newFish[age-1] += n
			}
		}
		fish = newFish
	}

	total := 0
	for _, n := range fish {
		total += n
	}
	println(total)
}
