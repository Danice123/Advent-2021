package helper

import (
	"os"
	"strconv"
	"strings"
)

func ReadInput(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

func ReadNumberList(file string) []int {
	data := ReadInput(file)[0]

	list := []int{}
	for _, s := range strings.Split(data, ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		list = append(list, i)
	}

	return list
}
