package day04

import (
	"regexp"
	"strconv"
	"strings"
)

var numberRegex = regexp.MustCompile(`(\d,[\d,]*)\n`)

type NumberIterator struct {
	numbers []int
	index   int
}

func (ths *NumberIterator) Next() int {
	val := ths.numbers[ths.index]
	ths.index++
	return val
}

func NewNumberIterator(data []byte) NumberIterator {
	numberData := numberRegex.FindSubmatch(data)
	if numberData == nil {
		panic("bad input")
	}

	iterator := NumberIterator{
		numbers: []int{},
	}

	numberList := strings.Split(string(numberData[1]), ",")
	for _, numberstring := range numberList {
		number, err := strconv.Atoi(numberstring)
		if err != nil {
			panic(err)
		}
		iterator.numbers = append(iterator.numbers, number)
	}

	return iterator
}
