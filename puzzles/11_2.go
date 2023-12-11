package puzzles

import (
	"arthur-fontaine/advent-of-code-2023/utils"
	"math"
	"os"
	"strings"
)

const collapsed_row_distance = 1000000

func parse_space_with_collapsed(space_string string) Space {
	space := Space{}

	rows := strings.Split(space_string, "\n")
	rotated_rows := strings.Split(utils.RotateString(space_string), "\n")

	for row_index, row := range rows {
		for column_index, character := range row {
			if string(character) == "#" {
				galaxy := Galaxy{
					row_index:    row_index,
					column_index: column_index,
				}
				space.galaxies = append(space.galaxies, galaxy)
			}
		}
	}

	for row_index, row := range rows {
		if utils.AllInArray(strings.Split(row, ""), ".") {
			space.collapsed_row_indexes = append(space.collapsed_row_indexes, row_index)
		}
	}

	for column_index, column := range rotated_rows {
		if utils.AllInArray(strings.Split(column, ""), ".") {
			space.collapsed_column_indexes = append(space.collapsed_column_indexes, column_index)
		}
	}

	return space
}

func get_galaxies_distance_by_collapses(space Space, galaxy_pair [2]Galaxy) [2]int {
	y_distance := 0
	min_row_index := int(math.Min(float64(galaxy_pair[0].row_index), float64(galaxy_pair[1].row_index)))
	max_row_index := int(math.Max(float64(galaxy_pair[0].row_index), float64(galaxy_pair[1].row_index)))

	for i := min_row_index; i < max_row_index; i++ {
		if utils.ArrayIncludes(space.collapsed_row_indexes, i) {
			y_distance += collapsed_row_distance
		} else {
			y_distance += 1
		}
	}

	x_distance := 0
	min_column_index := int(math.Min(float64(galaxy_pair[0].column_index), float64(galaxy_pair[1].column_index)))
	max_column_index := int(math.Max(float64(galaxy_pair[0].column_index), float64(galaxy_pair[1].column_index)))

	for i := min_column_index; i < max_column_index; i++ {
		if utils.ArrayIncludes(space.collapsed_column_indexes, i) {
			x_distance += collapsed_row_distance
		} else {
			x_distance += 1
		}
	}

	return [2]int{int(x_distance), int(y_distance)}
}

func day11_part2() any {
	input, err := os.ReadFile("resources/11/input.txt")
	if err != nil {
		panic(err)
	}
	input = []byte(strings.Trim(string(input), "\n"))

	space := parse_space_with_collapsed(string(input))

	galaxy_pairs := get_galaxy_pairs(space)

	sum_of_galaxies_distances := 0
	for _, galaxy_pair := range galaxy_pairs {
		galaxies_distance := get_galaxies_distance_by_collapses(space, galaxy_pair)
		x_distance := galaxies_distance[0]
		y_distance := galaxies_distance[1]

		sum_of_galaxies_distances += x_distance + y_distance
	}

	// We need to divide by 2 because the pair A-B also exists as B-A
	return sum_of_galaxies_distances / 2
}

func init() {
	RegisterPuzzle(11, 2, day11_part2)
}
