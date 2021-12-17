package day17

import (
	"advent2021/day05"
	"fmt"
	"math"
)

type Timing struct {
	velocity int
	time     int
	isZero   bool
}

func Part1() {

	goalXMax := 125
	goalXMin := 70

	validX := []Timing{}
	for xv := 0; xv <= goalXMax; xv++ {
		fx := 0
		for t := 0; true; t++ {
			xt := int(math.Max(0, float64(xv-t)))
			fx += xt
			if fx > goalXMax {
				break
			}
			if fx <= goalXMax && fx >= goalXMin {
				validX = append(validX, Timing{
					velocity: xv,
					time:     t,
				})
			}
			if xt == 0 {
				break
			}
		}
	}

	fmt.Printf("%v\n", validX)

	goalyMax := -121
	goalyMin := -159

	for yv := 0; yv < 200; yv++ {
		for _, timing := range validX {
			maxY := 0
			for t := timing.time; true; t++ {
				fy := 0
				for yt := 0; yt < t; yt++ {
					fy += yv - yt
					if fy > maxY {
						maxY = fy
					}
				}
				if fy < goalyMin {
					break
				}
				if fy >= goalyMin && fy <= goalyMax {
					fmt.Printf("x: %d, y: %d, t: %d, my: %d\n", timing.velocity, yv, timing.time, maxY)
				}
			}
		}
	}
}

func Part2() {
	goalXMax := 125
	goalXMin := 70

	validX := []Timing{}
	for xv := 0; xv <= goalXMax; xv++ {
		fx := 0
		for t := 0; true; t++ {
			xt := int(math.Max(0, float64(xv-t)))
			fx += xt
			if fx > goalXMax {
				break
			}
			if fx <= goalXMax && fx >= goalXMin {
				if xt == 0 {
					validX = append(validX, Timing{
						velocity: xv,
						time:     t,
						isZero:   true,
					})
				} else {
					validX = append(validX, Timing{
						velocity: xv,
						time:     t,
					})
				}
			}
			if xt == 0 {
				break
			}
		}
	}

	goalyMax := -121
	goalyMin := -159

	valid := map[day05.Coord]bool{}
	for yv := -400; yv < 400; yv++ {
		for _, timing := range validX {
			if timing.isZero {
				for t := timing.time; true; t++ {
					fy := 0
					for yt := 0; yt <= t; yt++ {
						fy += yv - yt
					}
					if fy < goalyMin {
						break
					}
					if fy >= goalyMin && fy <= goalyMax {
						valid[day05.Coord{
							X: timing.velocity,
							Y: yv,
						}] = true
						fmt.Printf("x: %d, y: %d, t: %d\n", timing.velocity, yv, timing.time)
					}
				}
			} else {
				fy := 0
				for yt := 0; yt <= timing.time; yt++ {
					fy += yv - yt
				}
				if fy >= goalyMin && fy <= goalyMax {
					valid[day05.Coord{
						X: timing.velocity,
						Y: yv,
					}] = true
					fmt.Printf("x: %d, y: %d, t: %d\n", timing.velocity, yv, timing.time)
				}
			}
		}
	}

	println(len(valid))
}
