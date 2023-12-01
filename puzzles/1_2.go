package puzzles

import (
	"arthur-fontaine/advent-of-code-2023/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var existing_digits = []struct {
	digit   int
	spelled string
}{
	{digit: 1, spelled: "one"},
	{digit: 2, spelled: "two"},
	{digit: 3, spelled: "three"},
	{digit: 4, spelled: "four"},
	{digit: 5, spelled: "five"},
	{digit: 6, spelled: "six"},
	{digit: 7, spelled: "seven"},
	{digit: 8, spelled: "eight"},
	{digit: 9, spelled: "nine"},
}

func get_digits(s string) []int {
	digits := []int{}

	for i, character := range s {
		if unicode.IsDigit(character) {
			digit, _ := strconv.Atoi(string(character))
			digits = append(digits, digit)
			continue
		}

		for _, existing_digit := range existing_digits {
			if utils.StringEndsWith(s[:i+1], existing_digit.spelled) {
				digits = append(digits, existing_digit.digit)
			}
		}
	}

	return digits
}

func day1_part2() any {
	input, err := os.ReadFile("resources/1/input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, line := range strings.Split(string(input), "\n") {
		digits := get_digits(line)
		if len(digits) == 0 {
			continue
		}

		number, _ := strconv.Atoi(fmt.Sprintf("%d%d", digits[0], digits[len(digits)-1]))
		sum += number
	}

	return sum
}

func init() {
	RegisterPuzzle(1, 2, day1_part2)
}
