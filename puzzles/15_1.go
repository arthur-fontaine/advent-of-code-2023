package puzzles

import (
	"os"
	"strings"
)

func get_ascii(character rune) int {
	return int(character)
}

func holiday_ash(s string) int {
	hash := 0

	for _, character := range s {
		ascii := get_ascii(character)
		hash += ascii
		hash *= 17
		hash %= 256
	}

	return hash
}

func day15_part1() any {
	input, err := os.ReadFile("resources/15/input.txt")
	if err != nil {
		panic(err)
	}
	input = []byte(strings.Trim(string(input), "\n"))

	initialization_sequence := strings.Split(string(input), ",")

	s := 0
	for _, step := range initialization_sequence {
		s += holiday_ash(step)
	}

	return s
}

func init() {
	RegisterPuzzle(15, 1, day15_part1)
}
