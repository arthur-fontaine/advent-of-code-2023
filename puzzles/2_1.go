package puzzles

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	color string
}

type GameSet struct {
	cubes []struct {
		cube  Cube
		count int
	}
}

type Game struct {
	id   int
	sets []GameSet
}

func parse_game_line(line string) (Game, error) {
	game := Game{}

	parts := strings.Split(line, " ")

	game_id_str := parts[1]
	parts = parts[2:]
	if !strings.HasSuffix(game_id_str, ":") {
		return game, fmt.Errorf("cannot find game id")
	}
	if game_id, err := strconv.Atoi(game_id_str[:len(game_id_str)-1]); err == nil {
		game.id = game_id
	} else {
		panic(err)
	}

	current_set := GameSet{}
	for i := 0; i < len(parts); i += 2 {
		cube_count_str := parts[i]
		cube_count, err := strconv.Atoi(cube_count_str)
		if err != nil {
			panic(err)
		}

		cube_color := parts[i+1]

		reset_current_set := false
		if strings.HasSuffix(cube_color, ";") {
			reset_current_set = true
		}

		if strings.HasSuffix(cube_color, ";") || strings.HasSuffix(cube_color, ",") {
			cube_color = cube_color[:len(cube_color)-1]
		}

		current_set.cubes = append(current_set.cubes, struct {
			cube  Cube
			count int
		}{
			cube:  Cube{color: cube_color},
			count: cube_count,
		})

		if reset_current_set {
			game.sets = append(game.sets, current_set)
			current_set = GameSet{}
		}
	}

	if len(current_set.cubes) > 0 {
		game.sets = append(game.sets, current_set)
	}

	return game, nil
}

func read_games(input string) []Game {
	lines := strings.Split(input, "\n")
	games := []Game{}

	for _, line := range lines {
		game, err := parse_game_line(line)
		if err != nil {
			panic(err)
		}
		games = append(games, game)
	}

	return games
}

func day2_part1() any {
	input, err := os.ReadFile("resources/2/input.txt")
	if err != nil {
		panic(err)
	}

	games := read_games(string((input)))
	available_cubes := []struct {
		cube  Cube
		count int
	}{
		{
			cube:  Cube{color: "red"},
			count: 12,
		},
		{
			cube:  Cube{color: "green"},
			count: 13,
		},
		{
			cube:  Cube{color: "blue"},
			count: 14,
		},
	}

	id_sum := 0
	for _, game := range games {
		game_is_possible := true

		for _, game_set := range game.sets {
			game_set_is_possible := true

			for _, cube := range game_set.cubes {

				// Find the cube according to color
				var found_available_cube struct {
					cube  Cube
					count int
				}
				for _, available_cube := range available_cubes {
					if available_cube.cube.color == cube.cube.color {
						found_available_cube = available_cube
						break
					}
				}
				if found_available_cube.cube.color == "" {
					continue
				}

				if found_available_cube.count < cube.count {
					fmt.Printf("Game is not possible because there is %d %s cubes, but only %d %s ones are available\n", cube.count, cube.cube.color, found_available_cube.count, found_available_cube.cube.color)
					fmt.Println(game)
					game_set_is_possible = false
					break
				}

			}

			if !game_set_is_possible {
				game_is_possible = false
				break
			}
		}

		if game_is_possible {
			id_sum += game.id
			fmt.Println("Game is possible", game)
		}
	}

	return id_sum
}

func init() {
	RegisterPuzzle(2, 1, day2_part1)
}
