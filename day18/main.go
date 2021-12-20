package day18

import (
	"advent2021/helper"
	"fmt"
	"strconv"
)

type Snailfish interface {
	Explode(depth int) bool
	Split() bool
	AddLeft(value int)
	AddRight(value int)

	Sum(Snailfish) Snailfish
	Magnitude() int

	IsPair() bool
	SetParent(parent *SnailfishPair)

	Sprint() string
	Copy(parent *SnailfishPair) Snailfish
}

type SnailfishPair struct {
	Parent *SnailfishPair
	Left   Snailfish
	Right  Snailfish
}

func (ths SnailfishPair) IsPair() bool {
	return true
}

func (ths *SnailfishPair) SetParent(parent *SnailfishPair) {
	ths.Parent = parent
}

func (ths *SnailfishPair) Explode(depth int) bool {
	if depth == 3 {
		if ths.Left.IsPair() {
			leftV := ths.Left.(*SnailfishPair).Left.(*SnailfishNumber).Value
			for current := ths; current.Parent != nil; current = current.Parent {
				if current.Parent.Left == current {
					continue
				} else {
					current.Parent.Left.AddRight(leftV)
					break
				}
			}

			rightV := ths.Left.(*SnailfishPair).Right.(*SnailfishNumber).Value
			ths.Right.AddLeft(rightV)

			ths.Left = &SnailfishNumber{}
			return true
		}
		if ths.Right.IsPair() {
			leftV := ths.Right.(*SnailfishPair).Left.(*SnailfishNumber).Value
			ths.Left.AddRight(leftV)

			rightV := ths.Right.(*SnailfishPair).Right.(*SnailfishNumber).Value
			for current := ths; current.Parent != nil; current = current.Parent {
				if current.Parent.Right == current {
					continue
				} else {
					current.Parent.Right.AddLeft(rightV)
					break
				}
			}
			ths.Right = &SnailfishNumber{}
			return true
		}
	} else {
		if ths.Left.IsPair() {
			if ths.Left.Explode(depth + 1) {
				return true
			}
		}
		if ths.Right.IsPair() {
			return ths.Right.Explode(depth + 1)
		}
	}
	return false
}

func (ths *SnailfishPair) AddLeft(value int) {
	ths.Left.AddLeft(value)
}

func (ths *SnailfishPair) AddRight(value int) {
	ths.Right.AddRight(value)
}

func (ths *SnailfishPair) Split() bool {
	if ths.Left.IsPair() {
		if ths.Left.Split() {
			return true
		}
	} else if ths.Left.(*SnailfishNumber).Value > 9 {
		v := ths.Left.(*SnailfishNumber).Value
		ths.Left = &SnailfishPair{
			Parent: ths,
			Left: &SnailfishNumber{
				Value: v / 2,
			},
			Right: &SnailfishNumber{
				Value: v/2 + v%2,
			},
		}
		return true
	}

	if ths.Right.IsPair() {
		return ths.Right.Split()
	} else if ths.Right.(*SnailfishNumber).Value > 9 {
		v := ths.Right.(*SnailfishNumber).Value
		ths.Right = &SnailfishPair{
			Parent: ths,
			Left: &SnailfishNumber{
				Value: v / 2,
			},
			Right: &SnailfishNumber{
				Value: v/2 + v%2,
			},
		}
		return true
	}

	return false
}

func (ths SnailfishPair) Sum(other Snailfish) Snailfish {
	sum := &SnailfishPair{}
	sum.Left = ths.Copy(sum)
	sum.Right = other.Copy(sum)

	for {
		if sum.Explode(0) {
			continue
		}
		if sum.Split() {
			continue
		}
		break
	}

	return sum
}

func (ths SnailfishPair) Magnitude() int {
	return ths.Left.Magnitude()*3 + ths.Right.Magnitude()*2
}

func (ths SnailfishPair) Sprint() string {
	return fmt.Sprintf("[%s,%s]", ths.Left.Sprint(), ths.Right.Sprint())
}

func (ths SnailfishPair) Copy(parent *SnailfishPair) Snailfish {
	new := &SnailfishPair{
		Parent: parent,
	}
	new.Left = ths.Left.Copy(new)
	new.Right = ths.Right.Copy(new)
	return new
}

type SnailfishNumber struct {
	Value int
}

func (ths SnailfishNumber) IsPair() bool {
	return false
}

func (ths *SnailfishNumber) SetParent(parent *SnailfishPair) {
}

func (ths *SnailfishNumber) Explode(depth int) bool {
	panic("Don't explode normal numbers")
}

func (ths *SnailfishNumber) AddLeft(value int) {
	ths.Value += value
}

func (ths *SnailfishNumber) AddRight(value int) {
	ths.Value += value
}

func (ths *SnailfishNumber) Split() bool {
	return false
}

func (ths SnailfishNumber) Sum(Snailfish) Snailfish {
	panic("Can't sum a normal number")
}

func (ths SnailfishNumber) Magnitude() int {
	return ths.Value
}

func (ths SnailfishNumber) Sprint() string {
	return fmt.Sprintf("%d", ths.Value)
}

func (ths SnailfishNumber) Copy(*SnailfishPair) Snailfish {
	return &SnailfishNumber{
		Value: ths.Value,
	}
}

func ParseNumber(input []byte, index int) (Snailfish, int) {
	if input[index] == '[' {
		number := &SnailfishPair{}
		number.Left, index = ParseNumber(input, index+1)
		number.Left.SetParent(number)
		if input[index] != ',' {
			panic("Missing comma at index " + strconv.Itoa(index))
		}
		number.Right, index = ParseNumber(input, index+1)
		number.Right.SetParent(number)
		if input[index] != ']' {
			panic("Missing close bracket at index " + strconv.Itoa(index))
		}
		return number, index + 1
	} else {
		return &SnailfishNumber{
			Value: int(input[index] - '0'),
		}, index + 1
	}
}

func Part1() {
	input := helper.ReadInput("day18/input.txt")

	numbers := []Snailfish{}
	for _, line := range input {
		n, _ := ParseNumber([]byte(line), 0)
		numbers = append(numbers, n)
	}

	sum := numbers[0]
	println(sum.Sprint())
	for i := 1; i < len(numbers); i++ {
		fmt.Printf("+ %s\n", numbers[i].Sprint())
		sum = sum.Sum(numbers[i])
		fmt.Printf("= %s\n\n", sum.Sprint())
	}

	println(sum.Magnitude())
}

func Part2() {
	input := helper.ReadInput("day18/input.txt")

	numbers := []Snailfish{}
	for _, line := range input {
		n, _ := ParseNumber([]byte(line), 0)
		numbers = append(numbers, n)
	}

	biggest := 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			}
			sum := numbers[i].Sum(numbers[j])
			mag := sum.Magnitude()
			if mag > biggest {
				biggest = mag
			}
		}
	}

	println(biggest)
}
