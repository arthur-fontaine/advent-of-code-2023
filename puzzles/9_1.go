package puzzles

import (
	"arthur-fontaine/advent-of-code-2023/utils"
	"os"
	"strconv"
	"strings"
)

func row_to_number_list(row string) []int {
	numbers_str := strings.Split(string(row), " ")
	number_list := []int{}

	for _, number_str := range numbers_str {
		if number_str == "" {
			continue
		}

		number, err := strconv.Atoi(number_str)
		if err != nil {
			panic(err)
		}
		number_list = append(number_list, number)
	}

	return number_list
}

func get_subline(numbers []int) []int {
	subline := make([]int, len(numbers)-1)

	for i := range numbers {
		if i == len(numbers)-1 {
			continue
		}

		subline[i] = numbers[i+1] - numbers[i]
	}

	return subline
}

func get_pyramid(numbers []int) [][]int {
	pyramid := [][]int{numbers}

	for zero_line_reached := false; !zero_line_reached; zero_line_reached = utils.AllInArray(pyramid[len(pyramid)-1], 0) {
		pyramid = append(pyramid, get_subline(pyramid[len(pyramid)-1]))
	}

	return pyramid
}

func ascend_pyramid(pyramid [][]int) [][]int {
	for i := len(pyramid) - 1; i > 0; i-- {
		if i == len(pyramid)-1 {
			pyramid[i] = append(pyramid[i], 0)
		}

		line := pyramid[i]
		above_line := pyramid[i-1]

		last_element := line[len(line)-1]
		above_last_element := above_line[len(above_line)-1]

		pyramid[i-1] = append(pyramid[i-1], above_last_element+last_element)
	}

	return pyramid
}

func day9_part1() any {
	input, err := os.ReadFile("resources/9/input.txt")
	if err != nil {
		panic(err)
	}
	input = []byte(strings.Trim(string(input), "\n"))

	rows := strings.Split(string(input), "\n")

	sum_of_extrapolated_values := 0
	for _, row := range rows {
		pyramid := get_pyramid(row_to_number_list(row))
		extrapolated_pyramid := ascend_pyramid(pyramid)

		sum_of_extrapolated_values += extrapolated_pyramid[0][len(extrapolated_pyramid[0])-1]
	}

	return sum_of_extrapolated_values
}

func init() {
	RegisterPuzzle(9, 1, day9_part1)
}
