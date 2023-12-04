package puzzles

import (
	"math"
	"os"
	"strings"
)

func get_points(card string) int {
	numbers := strings.Split(card, ":")[1]

	separated_numbers := strings.Split(numbers, "|")
	winning_numbers := strings.Split(strings.Trim(separated_numbers[0], " "), " ")
	own_numbers := strings.Split(strings.Trim(separated_numbers[1], " "), " ")

	power := -1
	for _, own_number := range own_numbers {
		for _, winning_number := range winning_numbers {
			if len(own_number) == 0 || len(winning_number) == 0 {
				continue
			}

			if own_number == winning_number {
				power++
				break
			}
		}
	}

	if power == -1 {
		return 0
	}

	return int(math.Pow(2, float64(power)))
}

func day4_part1() any {
	input, err := os.ReadFile("resources/4/input.txt")
	if err != nil {
		panic(err)
	}

	sum_point := 0
	for _, card_line := range strings.Split(string(input), "\n") {
		sum_point += get_points(card_line)
	}

	return sum_point
}

func init() {
	RegisterPuzzle(4, 1, day4_part1)
}
