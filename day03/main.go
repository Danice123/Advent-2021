package day03

import (
	"advent2021/helper"
	"fmt"
)

// https://stackoverflow.com/questions/23192262/how-would-you-set-and-clear-a-single-bit-in-go

func clearBit(n uint, pos int) uint {
	mask := ^(uint(1) << pos)
	n &= mask
	return n
}

func setBit(n uint, pos int) uint {
	n |= (1 << pos)
	return n
}

func findSumsOfDataList(data []string) []int {
	sums := make([]int, len(data[0]))
	for _, bs := range data {
		for i, b := range bs {
			if b == '1' {
				sums[i]++
			}
		}
	}
	return sums
}

func Part1() {
	data := helper.ReadInput("day03/input.txt")
	sums := findSumsOfDataList(data)

	var gamma uint
	for i, sum := range sums {
		if sum > len(data)/2 {
			gamma = setBit(gamma, len(sums)-i-1)
		}
	}
	fmt.Printf("%064b\n", gamma)

	epsilon := ^gamma
	for i := len(sums); i < 64; i++ {
		epsilon = clearBit(epsilon, i)
	}
	fmt.Printf("%064b\n", epsilon)

	println(gamma * epsilon)
}

func recurseToFindValueForOxygen(input []string, index int) []string {
	sums := findSumsOfDataList(input)

	newInput := []string{}
	for _, bs := range input {
		ones := sums[index]
		zeros := len(input) - sums[index]

		if (ones >= zeros && bs[index] == '1') || (ones < zeros && bs[index] == '0') {
			newInput = append(newInput, bs)
		}
	}
	fmt.Printf("%d: %v\n", index, newInput)

	if len(newInput) > 1 {
		newInput = recurseToFindValueForOxygen(newInput, index+1)
	}
	return newInput
}

func recurseToFindValueForC02(input []string, index int) []string {
	sums := findSumsOfDataList(input)

	newInput := []string{}
	for _, bs := range input {
		ones := sums[index]
		zeros := len(input) - sums[index]

		if (ones >= zeros && bs[index] == '0') || (ones < zeros && bs[index] == '1') {
			newInput = append(newInput, bs)
		}
	}
	fmt.Printf("%d: %v\n", index, newInput)

	if len(newInput) > 1 {
		newInput = recurseToFindValueForC02(newInput, index+1)
	}
	return newInput
}

func byteStringToInt(bs string) uint {
	var value uint
	for i, bit := range bs {
		if bit == '1' {
			value = setBit(value, len(bs)-i-1)
		}
	}
	return value
}

func Part2() {
	data := helper.ReadInput("day03/input.txt")

	oxygen := recurseToFindValueForOxygen(data, 0)[0]
	println(oxygen)

	c02 := recurseToFindValueForC02(data, 0)[0]
	println(c02)

	println(byteStringToInt(oxygen) * byteStringToInt(c02))
}
