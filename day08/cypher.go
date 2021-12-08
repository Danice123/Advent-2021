package day08

import (
	"fmt"
	"sort"
	"strings"
)

type Cypher struct {
	cypher map[rune]map[rune]int
}

func NewCypher() *Cypher {
	return &Cypher{
		cypher: map[rune]map[rune]int{
			'a': {'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0},
			'b': {'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0},
			'c': {'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0},
			'd': {'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0},
			'e': {'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0},
			'f': {'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0},
			'g': {'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0},
		},
	}
}

func (ths Cypher) Has(c rune, v rune) bool {
	_, ok := ths.cypher[c][v]
	return ok
}

func (ths *Cypher) NarrowPossibleOptions(c rune, possible map[rune]int) {
	for option := range ths.cypher[c] {
		if _, ok := possible[option]; !ok {
			delete(ths.cypher[c], option)
		}
	}

	for c, options := range ths.cypher {
		if len(options) == 1 {
			var takenC rune
			for r := range options {
				takenC = r
				break
			}
			for checkC, checkOptions := range ths.cypher {
				if checkC == c {
					continue
				}
				if _, ok := checkOptions[takenC]; ok {
					delete(ths.cypher[checkC], takenC)
				}
			}
		}
	}
}

func (ths Cypher) IsSolved() bool {
	for _, options := range ths.cypher {
		if len(options) > 1 {
			return false
		}
	}
	return true
}

func (ths Cypher) Solve(output []string) string {
	reverseCypher := make(map[rune]rune)
	for c, options := range ths.cypher {
		for v := range options {
			reverseCypher[v] = c
			break
		}
	}

	solution := ""
	for _, s := range output {
		uncyphered := []string{}
		for _, r := range s {
			uncyphered = append(uncyphered, string([]byte{byte(reverseCypher[r])}))
		}
		sort.Strings(uncyphered)
		switch strings.Join(uncyphered, "") {
		case "abcefg":
			solution += "0"
		case "cf":
			solution += "1"
		case "acdeg":
			solution += "2"
		case "acdfg":
			solution += "3"
		case "bcdf":
			solution += "4"
		case "abdfg":
			solution += "5"
		case "abdefg":
			solution += "6"
		case "acf":
			solution += "7"
		case "abcdefg":
			solution += "8"
		case "abcdfg":
			solution += "9"
		default:
			solution += "?"
		}
	}

	return solution
}

func (ths Cypher) Print() {
	l := []string{}
	for c := range ths.cypher {
		l = append(l, string([]byte{byte(c)}))
	}
	sort.Strings(l)

	for _, c := range l {
		fmt.Printf("%s: ", c)
		for option := range ths.cypher[rune(c[0])] {
			fmt.Printf("%c, ", option)
		}
		fmt.Print("\n")
	}
}
