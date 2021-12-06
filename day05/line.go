package day05

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Coord struct {
	X int
	Y int
}

func (ths Coord) String() string {
	return fmt.Sprintf("%d,%d", ths.X, ths.Y)
}

func NewCoord(input string) Coord {
	numbers := strings.Split(input, ",")
	x, err := strconv.Atoi(numbers[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(numbers[1])
	if err != nil {
		panic(err)
	}
	return Coord{
		X: x,
		Y: y,
	}
}

type Line struct {
	Start Coord
	End   Coord
}

func (ths Line) String() string {
	return fmt.Sprintf("%s => %s", ths.Start.String(), ths.End.String())
}

func (ths Line) GetCoverage() []Coord {
	coverage := []Coord{}

	if ths.Start.X == ths.End.X {
		diff := int(math.Abs(float64(ths.End.Y - ths.Start.Y)))
		for i := 0; i <= diff; i++ {
			var c Coord
			if ths.Start.Y > ths.End.Y {
				c = Coord{
					X: ths.Start.X,
					Y: ths.Start.Y - i,
				}
			} else {
				c = Coord{
					X: ths.Start.X,
					Y: ths.Start.Y + i,
				}
			}
			coverage = append(coverage, c)
		}
	} else if ths.Start.Y == ths.End.Y {
		diff := int(math.Abs(float64(ths.End.X - ths.Start.X)))
		for i := 0; i <= diff; i++ {
			var c Coord
			if ths.Start.X > ths.End.X {
				c = Coord{
					X: ths.Start.X - i,
					Y: ths.Start.Y,
				}
			} else {
				c = Coord{
					X: ths.Start.X + i,
					Y: ths.Start.Y,
				}
			}
			coverage = append(coverage, c)
		}
	} else {
		x := ths.Start.X
		y := ths.Start.Y
		for x != ths.End.X && y != ths.End.Y {
			coverage = append(coverage, Coord{
				X: x,
				Y: y,
			})

			if ths.Start.X > ths.End.X {
				x--

			} else {
				x++
			}
			if ths.Start.Y > ths.End.Y {
				y--
			} else {
				y++
			}
		}
		coverage = append(coverage, Coord{
			X: ths.End.X,
			Y: ths.End.Y,
		})
	}

	return coverage
}

func NewLine(input [][]byte) Line {
	return Line{
		Start: NewCoord(string(input[1])),
		End:   NewCoord(string(input[2])),
	}
}
