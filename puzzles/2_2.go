package puzzles

import (
	"os"
)

func get_minimum_set(game Game) []struct {
	cube  Cube
	count int
} {
	minimum_set := []struct {
		cube  Cube
		count int
	}{}

	for _, game_set := range game.sets {
		for _, cube := range game_set.cubes {
			var found_listed_cube struct {
				cube  Cube
				count int
			}
			for i, listed_cube := range minimum_set {
				if listed_cube.cube.color == cube.cube.color {
					found_listed_cube = listed_cube
					minimum_set = append(minimum_set[0:i], minimum_set[i+1:]...)
					break
				}
			}

			if found_listed_cube.cube.color == "" {
				minimum_set = append(minimum_set, cube)
			} else if found_listed_cube.count > cube.count {
				minimum_set = append(minimum_set, found_listed_cube)
			} else {
				minimum_set = append(minimum_set, cube)
			}
		}
	}

	return minimum_set
}

func get_set_power(set []struct {
	cube  Cube
	count int
}) int {
	if len(set) == 0 {
		return 0
	}

	set_power := 1

	for _, cube := range set {
		set_power *= cube.count
	}

	return set_power
}

func day2_part2() any {
	input, err := os.ReadFile("resources/2/input.txt")
	if err != nil {
		panic(err)
	}

	games := read_games(string((input)))
	power_sum := 0
	for _, game := range games {
		power_sum += get_set_power(get_minimum_set(game))
	}

	return power_sum
}

func init() {
	RegisterPuzzle(2, 2, day2_part2)
}
