package puzzles

import (
	"os"
	"strings"
)

func ascend_pyramid_left(pyramid [][]int) [][]int {
	for i := len(pyramid) - 1; i > 0; i-- {
		if i == len(pyramid)-1 {
			pyramid[i] = append(pyramid[i], 0)
		}

		line := pyramid[i]
		above_line := pyramid[i-1]

		first_element := line[0]
		above_first_element := above_line[0]

		pyramid[i-1] = append([]int{above_first_element - first_element}, pyramid[i-1]...)
	}

	return pyramid
}

func day9_part2() any {
	input, err := os.ReadFile("resources/9/input.txt")
	if err != nil {
		panic(err)
	}
	input = []byte(strings.Trim(string(input), "\n"))

	rows := strings.Split(string(input), "\n")

	sum_of_extrapolated_values := 0
	for _, row := range rows {
		pyramid := get_pyramid(row_to_number_list(row))
		extrapolated_pyramid_by_left := ascend_pyramid_left(pyramid)

		sum_of_extrapolated_values += extrapolated_pyramid_by_left[0][0]
	}

	return sum_of_extrapolated_values
}

func init() {
	RegisterPuzzle(9, 2, day9_part2)
}
