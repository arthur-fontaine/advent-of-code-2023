package puzzles

import (
	. "arthur-fontaine/advent-of-code-2023/utils"
	"fmt"
	"os"
	"strings"
)

func find_smudge_index_in_line(line_a string, line_b string) int {
	differences := []int{}
	for character_a_index, character_a := range line_a {
		character_b := rune(line_b[character_a_index])
		if character_a != character_b {
			differences = append(differences, character_a_index)
		}
	}

	if len(differences) != 1 {
		return -1
	}

	return differences[0]
}

func check_reflection_with_smudge(strings_a []string, strings_b []string) bool {
	max_length := min(len(strings_a), len(strings_b))

	// Pick the max_length elements from each array
	a := strings_a[len(strings_a)-max_length:]
	b := ReverseArray(strings_b)[len(strings_b)-max_length:]

	found_smudge := false
	for line_a_index, line_a := range a {
		line_b := b[line_a_index]
		if smudge_index := find_smudge_index_in_line(line_a, line_b); smudge_index != -1 {
			var replaced_smudge string
			if a[line_a_index][smudge_index] == '.' {
				replaced_smudge = "#"
			} else {
				replaced_smudge = "."
			}

			a[line_a_index] = a[line_a_index][:smudge_index] + replaced_smudge + a[line_a_index][smudge_index+1:]
			found_smudge = true
			break
		}
	}

	return found_smudge && ArraysAreSame(a, b)
}

func get_vertical_line_index_reflexion_with_smudge(pattern string) int {
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

		if check_reflection_with_smudge(splitted_rows[0], splitted_rows[1]) {
			return i - 1
		}
	}

	return -1
}

func get_horizontal_line_index_reflexion_with_smudge(pattern string) int {
	return get_vertical_line_index_reflexion_with_smudge(RotateString(pattern))
}

func day13_part2() any {
	input, err := os.ReadFile("resources/13/input.txt")
	if err != nil {
		panic(err)
	}
	input = []byte(strings.Trim(string(input), "\n"))

	patterns := read_patterns(string(input))
	s := 0
	for i, pattern := range patterns {
		v := get_vertical_line_index_reflexion_with_smudge(pattern)
		if v != -1 {
			s += (v + 1) * 100
			continue
		}

		h := get_horizontal_line_index_reflexion_with_smudge(pattern)
		if h != -1 {
			s += h + 1
			continue
		}

		panic(fmt.Sprintf("No reflexion found for following pattern (index %d):\n%s", i, pattern))
	}

	return s
}

func init() {
	RegisterPuzzle(13, 2, day13_part2)
}
