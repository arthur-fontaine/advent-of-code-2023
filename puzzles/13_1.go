package puzzles

import (
	. "arthur-fontaine/advent-of-code-2023/utils"
	"fmt"
	"os"
	"strings"
)

func read_patterns(input string) []string {
	return strings.Split(input, "\n\n")
}

func check_reflection(strings_a []string, strings_b []string) bool {
	max_length := min(len(strings_a), len(strings_b))

	// Pick the max_length elements from each array
	a := strings_a[len(strings_a)-max_length:]
	b := ReverseArray(strings_b)[len(strings_b)-max_length:]

	return ArraysAreSame(a, b)
}

func get_vertical_line_index_reflexion(pattern string) int {
	rows := strings.Split(pattern, "\n")

	for i := 1; i < len(rows); i++ {
		splitted_rows := SplitArrayAt(rows, i)

		max_length := min(len(splitted_rows[0]), len(splitted_rows[1]))

		if len(splitted_rows[0]) > max_length {
			splitted_rows[0] = splitted_rows[0][len(splitted_rows[0])-max_length:]
		} else {
			splitted_rows[0] = splitted_rows[0][:max_length]
		}

		splitted_rows[1] = splitted_rows[1][:max_length]

		if check_reflection(splitted_rows[0], splitted_rows[1]) {
			return i - 1
		}
	}

	return -1
}

func get_horizontal_line_index_reflexion(pattern string) int {
	return get_vertical_line_index_reflexion(RotateString(pattern))
}

func day13_part1() any {
	input, err := os.ReadFile("resources/13/input.txt")
	if err != nil {
		panic(err)
	}
	input = []byte(strings.Trim(string(input), "\n"))

	patterns := read_patterns(string(input))
	s := 0
	for i, pattern := range patterns {
		v := get_vertical_line_index_reflexion(pattern)
		if v != -1 {
			s += (v + 1) * 100
			continue
		}

		h := get_horizontal_line_index_reflexion(pattern)
		if h != -1 {
			s += h + 1
			continue
		}

		panic(fmt.Sprintf("No reflexion found for following pattern (index %d):\n%s", i, pattern))
	}

	return s
}

func init() {
	RegisterPuzzle(13, 1, day13_part1)
}
