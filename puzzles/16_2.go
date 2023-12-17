package puzzles

import (
	"os"
	"strings"
)

func find_all_initial_beam_configurations(layout []string) [][2][2]int {
	all_initial_beam_configurations := [][2][2]int{}

	for i := range layout {
		// Handle start left and right
		all_initial_beam_configurations = append(all_initial_beam_configurations, [2][2]int{{-1, i}, {1, 0}})
		all_initial_beam_configurations = append(all_initial_beam_configurations, [2][2]int{{len(layout[0]), i}, {-1, 0}})
	}

	for i := range layout[0] {
		// Handle start bottom and top
		all_initial_beam_configurations = append(all_initial_beam_configurations, [2][2]int{{i, -1}, {0, 1}})
		all_initial_beam_configurations = append(all_initial_beam_configurations, [2][2]int{{i, len(layout)}, {0, -1}})
	}

	return all_initial_beam_configurations
}

func day16_part2() any {
	input, err := os.ReadFile("resources/16/input.txt")
	if err != nil {
		panic(err)
	}
	input = []byte(strings.Trim(string(input), "\n"))

	layout := strings.Split(string(input), "\n")

	energized_tiles := 0
	for _, initial_beam_configuration := range find_all_initial_beam_configurations(layout) {
		beam_layout := make([]string, len(layout))
		copy(beam_layout, layout)
		beam_layout = run_beam(initial_beam_configuration[0], initial_beam_configuration[1], layout, beam_layout, map[[2][2]int]bool{})
		current_energized_tiles := strings.Count(strings.Join(beam_layout, ""), "#")
		if current_energized_tiles > energized_tiles {
			energized_tiles = current_energized_tiles
		}
	}

	return energized_tiles
}

func init() {
	RegisterPuzzle(16, 2, day16_part2)
}
