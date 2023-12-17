package puzzles

import (
	"os"
	"strings"
)

func run_beam(position [2]int, movement [2]int, layout []string, beam_layout []string, history map[[2][2]int]bool) []string {
	if _, ok := history[[2][2]int{position, movement}]; ok {
		return beam_layout
	}

	history[[2][2]int{position, movement}] = true

	// Validate movement
	if (movement[0] != 0 && movement[1] != 0) || (movement[0] == movement[1]) {
		panic("unexepected movement")
	}

	new_position := [2]int{position[0] + movement[0], position[1] + movement[1]}

	if new_position[0] < 0 || new_position[0] > len(layout[0])-1 || new_position[1] < 0 || new_position[1] > len(layout)-1 {
		return beam_layout
	}

	char := layout[new_position[1]][new_position[0]]

	beam_layout[new_position[1]] = beam_layout[new_position[1]][:new_position[0]] + "#" + beam_layout[new_position[1]][new_position[0]+1:]

	switch char {
	case '.':
		beam_layout = run_beam(new_position, movement, layout, beam_layout, history)
	case '|':
		{
			if movement[1] != 0 {
				// If the beam encounters the pointy end of a splitter
				beam_layout = run_beam(new_position, movement, layout, beam_layout, history)
			} else {
				// If the beam encounters the flat side of a splitter
				beam_layout = run_beam(new_position, [2]int{0, -1}, layout, beam_layout, history)
				beam_layout = run_beam(new_position, [2]int{0, 1}, layout, beam_layout, history)
			}
		}
	case '-':
		{
			if movement[0] != 0 {
				// If the beam encounters the pointy end of a splitter
				beam_layout = run_beam(new_position, movement, layout, beam_layout, history)
			} else {
				// If the beam encounters the flat side of a splitter
				beam_layout = run_beam(new_position, [2]int{-1, 0}, layout, beam_layout, history)
				beam_layout = run_beam(new_position, [2]int{1, 0}, layout, beam_layout, history)
			}
		}
	case '/':
		{
			if movement[0] == -1 {
				beam_layout = run_beam(new_position, [2]int{0, 1}, layout, beam_layout, history)
			} else if movement[0] == 1 {
				beam_layout = run_beam(new_position, [2]int{0, -1}, layout, beam_layout, history)
			} else if movement[1] == -1 {
				beam_layout = run_beam(new_position, [2]int{1, 0}, layout, beam_layout, history)
			} else if movement[1] == 1 {
				beam_layout = run_beam(new_position, [2]int{-1, 0}, layout, beam_layout, history)
			}
		}
	case '\\':
		{
			if movement[0] == -1 {
				beam_layout = run_beam(new_position, [2]int{0, -1}, layout, beam_layout, history)
			} else if movement[0] == 1 {
				beam_layout = run_beam(new_position, [2]int{0, 1}, layout, beam_layout, history)
			} else if movement[1] == -1 {
				beam_layout = run_beam(new_position, [2]int{-1, 0}, layout, beam_layout, history)
			} else if movement[1] == 1 {
				beam_layout = run_beam(new_position, [2]int{1, 0}, layout, beam_layout, history)
			}
		}
	}

	return beam_layout
}

func day16_part1() any {
	input, err := os.ReadFile("resources/16/input.txt")
	if err != nil {
		panic(err)
	}
	input = []byte(strings.Trim(string(input), "\n"))

	layout := strings.Split(string(input), "\n")
	beam_layout := make([]string, len(layout))
	copy(beam_layout, layout)

	beam_layout = run_beam([2]int{-1, 0}, [2]int{1, 0}, layout, beam_layout, map[[2][2]int]bool{})

	return strings.Count(strings.Join(beam_layout, ""), "#")
}

func init() {
	RegisterPuzzle(16, 1, day16_part1)
}
