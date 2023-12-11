package puzzles

import (
	"arthur-fontaine/advent-of-code-2023/utils"
	"math"
	"os"
	"strings"
)

func horizontally_normalize_space(space_string string) string {
	rows := strings.Split(space_string, "\n")
	new_rows := []string{}

	for _, row := range rows {
		new_rows = append(new_rows, row)

		if utils.AllInArray(strings.Split(row, ""), ".") {
			// Add the line twice if the row is full of dot
			new_rows = append(new_rows, row)
		}
	}

	return strings.Join(new_rows, "\n")
}

func vertically_normalize_space(space_string string) string {
	rotated_space := utils.RotateString(space_string)
	rotated_space = horizontally_normalize_space(rotated_space)
	space_string = utils.RotateString(rotated_space)

	return space_string
}

func normalize_space(space_string string) string {
	space_string = horizontally_normalize_space(space_string)
	space_string = vertically_normalize_space(space_string)

	return space_string
}

type Space struct {
	galaxies []Galaxy
}

type Galaxy struct {
	row_index    int
	column_index int
}

func parse_space(space_string string) Space {
	space := Space{}
	galaxies := []Galaxy{}

	rows := strings.Split(space_string, "\n")

	for row_index, row := range rows {
		for column_index, character := range row {
			if string(character) == "#" {
				galaxy := Galaxy{
					row_index:    row_index,
					column_index: column_index,
				}
				galaxies = append(galaxies, galaxy)
			}
		}
	}

	space.galaxies = galaxies

	return space
}

func get_galaxy_pairs(space Space) [][2]Galaxy {
	galaxy_pairs := [][2]Galaxy{}

	for _, galaxy_a := range space.galaxies {
		for _, galaxy_b := range space.galaxies {
			if galaxy_a == galaxy_b {
				continue
			}

			galaxy_pairs = append(galaxy_pairs, [2]Galaxy{galaxy_a, galaxy_b})
		}
	}

	return galaxy_pairs
}

func get_galaxies_distance(galaxy_pair [2]Galaxy) [2]int {
	x_distance := math.Abs(float64(galaxy_pair[0].column_index - galaxy_pair[1].column_index))
	y_distance := math.Abs(float64(galaxy_pair[0].row_index - galaxy_pair[1].row_index))

	return [2]int{int(x_distance), int(y_distance)}
}

func day11_part1() any {
	input, err := os.ReadFile("resources/11/input.txt")
	if err != nil {
		panic(err)
	}
	input = []byte(strings.Trim(string(input), "\n"))

	space_string := normalize_space(string(input))
	space := parse_space(space_string)

	galaxy_pairs := get_galaxy_pairs(space)

	sum_of_galaxies_distances := 0
	for _, galaxy_pair := range galaxy_pairs {
		galaxies_distance := get_galaxies_distance(galaxy_pair)
		x_distance := galaxies_distance[0]
		y_distance := galaxies_distance[1]

		sum_of_galaxies_distances += x_distance + y_distance
	}

	// We need to divide by 2 because the pair A-B also exists as B-A
	return sum_of_galaxies_distances / 2
}

func init() {
	RegisterPuzzle(11, 1, day11_part1)
}
