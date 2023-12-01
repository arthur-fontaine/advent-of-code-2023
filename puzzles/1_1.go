package puzzles

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func get_first_digit(s string) (int, bool) {
	for _, character := range s {
		if unicode.IsDigit(character) {
			number, err := strconv.Atoi(string(character))
			if err != nil {
				panic(err)
			}

			return number, true
		}
	}

	return 0, false
}

func reverse(s string) string {
	reversed_string := ""
	for _, character := range s {
		reversed_string = string(character) + reversed_string
	}
	return reversed_string
}

func get_last_digit(s string) (int, bool) {
	return get_first_digit(reverse(s))
}

func trebuchet() any {
	input, err := os.ReadFile("resources/1_1/input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, line := range strings.Split(string(input), "\n") {
		first_digit, _ := get_first_digit(line)
		last_digit, _ := get_last_digit(line)

		number, err := strconv.Atoi(fmt.Sprintf("%d%d", first_digit, last_digit))
		if err != nil {
			panic(err)
		}

		sum += number
	}

	return sum
}

func init() {
	RegisterPuzzle(1, 1, trebuchet)
}
