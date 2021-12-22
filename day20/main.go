package day20

import (
	"advent2021/day05"
	"advent2021/helper"
	"fmt"
)

type Coord struct {
	X int
	Y int
}

type Image struct {
	Values   map[Coord]bool
	Min      Coord
	Max      Coord
	Infinity bool
}

func (ths Image) GetPixelValue(center day05.Coord) int {
	index := 0
	val := 0
	for y := center.Y - 1; y < center.Y+2; y++ {
		for x := center.X - 1; x < center.X+2; x++ {
			index++
			if v, ok := ths.Values[Coord{X: x, Y: y}]; ok {
				if v {
					val |= (1 << (9 - index))
				}
			} else {
				if ths.Infinity {
					val |= (1 << (9 - index))
				}
			}
		}
	}
	return val
}

func (ths Image) GenerateOutput(enhancement []byte) Image {
	expanded := 1
	output := Image{
		Values: map[Coord]bool{},
		Min: Coord{
			X: ths.Min.X - expanded,
			Y: ths.Min.Y - expanded,
		},
		Max: Coord{
			X: ths.Max.X + expanded,
			Y: ths.Max.Y + expanded,
		},
	}

	if ths.Infinity {
		output.Infinity = enhancement[len(enhancement)-1] == '#'
	} else {
		output.Infinity = enhancement[0] == '#'
	}

	for y := output.Min.Y; y < output.Max.Y; y++ {
		for x := output.Min.X; x < output.Max.X; x++ {
			val := ths.GetPixelValue(day05.Coord{
				X: x,
				Y: y,
			})
			output.Values[Coord{X: x, Y: y}] = enhancement[val] == '#'
		}
	}
	return output
}

func (ths Image) NumLit() int {
	sum := 0
	for y := ths.Min.Y; y < ths.Max.Y; y++ {
		for x := ths.Min.X; x < ths.Max.X; x++ {
			if ths.Values[Coord{X: x, Y: y}] {
				sum++
			}
		}
	}
	return sum
}

func (ths Image) Print() {
	for y := ths.Min.Y; y < ths.Max.Y; y++ {
		for x := ths.Min.X; x < ths.Max.X; x++ {
			if ths.Values[Coord{X: x, Y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func NewImage(input []string) (Image, []byte) {
	var enhance []byte
	image := Image{
		Values: map[Coord]bool{},
		Min: Coord{
			X: 0,
			Y: 0,
		},
		Max: Coord{},
	}
	y := 0
	for i, line := range input {
		if line == "" {
			continue
		}
		if i == 0 {
			enhance = []byte(line)
			continue
		}

		image.Max.X = len(line)
		for x, r := range line {
			image.Values[Coord{X: x, Y: y}] = r == '#'
		}
		y++
	}
	image.Max.Y = y
	return image, enhance
}

func Part1() {
	input := helper.ReadInput("day20/input.txt")
	image, enhance := NewImage(input)

	image.Print()
	newImage := image.GenerateOutput(enhance)
	newImage.Print()
	newImage = newImage.GenerateOutput(enhance)
	newImage.Print()
	println(newImage.NumLit())
}

func Part2() {
	input := helper.ReadInput("day20/input.txt")
	image, enhance := NewImage(input)

	for i := 0; i < 50; i++ {
		image = image.GenerateOutput(enhance)
		// image.Print()
	}

	println(image.NumLit())
}
