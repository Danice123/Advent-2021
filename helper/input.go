package helper

import (
	"os"
	"strings"
)

func ReadInput(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}
