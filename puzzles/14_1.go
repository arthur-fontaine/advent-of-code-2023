package puzzles

import (
	"os"
	"strings"
)

func up_all_rounded_rock(platform string) string {
	rows := strings.Split(platform, "\n")

	for i := len(rows) - 1; i > 0; i-- {
		row := &rows[i]
		upper_row := &rows[i-1]

		old_row := *row

		for j := 0; j < len(*row); j++ {
			character := (*row)[j]
			upper_character := (*upper_row)[j]

			if character == 'O' && upper_character == '.' {
				*row = (*row)[:j] + "." + (*row)[j+1:]
				*upper_row = (*upper_row)[:j] + "O" + (*upper_row)[j+1:]
			}
		}

		if old_row != *row {
			i = min(i+2, len(rows))
		}
	}

	return strings.Join(rows, "\n")
}

func get_platform_score(platform string) int {
	score := 0

	rows := strings.Split(platform, "\n")

	for i, v := range rows {
		score += strings.Count(v, "O") * (len(rows) - i)
	}

	return score
}

func day14_part1() any {
	input, err := os.ReadFile("resources/14/input.txt")
	if err != nil {
		panic(err)
	}
	input = []byte(strings.Trim(string(input), "\n"))

	return get_platform_score(up_all_rounded_rock(string(input)))
}

func init() {
	RegisterPuzzle(14, 1, day14_part1)
}
