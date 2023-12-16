package puzzles

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func remove_label_from_box(box *[][2]string, label string) {
	for i, lens := range *box {
		if lens[0] == label {
			*box = append((*box)[:i], (*box)[i+1:]...)
		}
	}
}

func replace_or_push_to_box(box *[][2]string, lens_to_add [2]string) {
	for i, lens := range *box {
		if lens[0] == lens_to_add[0] {
			(*box)[i] = lens_to_add
			return
		}
	}

	*box = append(*box, lens_to_add)
}

func parse_step(step string) ([3]string, error) {
	if strings.Count(step, "-") > 0 {
		step_parts := strings.Split(step, "-")
		return [3]string{step_parts[0], "-", step_parts[1]}, nil
	}

	if strings.Count(step, "=") > 0 {
		step_parts := strings.Split(step, "=")
		return [3]string{step_parts[0], "=", step_parts[1]}, nil
	}

	return [3]string{}, fmt.Errorf("unable to parse step %v", step)
}

func place_lens(boxes *[256][][2]string, step_str string) {
	step, err := parse_step(step_str)
	if err != nil {
		panic(err)
	}

	label := step[0]
	box_index := holiday_ash(label)
	focal_length := step[2]

	lens := [2]string{label, focal_length}

	switch step[1] {
	case "=":
		replace_or_push_to_box(&(*boxes)[box_index], lens)
	case "-":
		remove_label_from_box(&(*boxes)[box_index], label)
	}
}

func place_lenses(boxes *[256][][2]string, steps []string) {
	for _, step := range steps {
		place_lens(boxes, step)
	}
}

func calculate_focusing_power(boxes [256][][2]string) int {
	focusing_power := 0

	for box_number, box := range boxes {
		for lens_index, lens := range box {
			focal_length, err := strconv.Atoi(lens[1])
			if err != nil {
				panic(err)
			}
			focusing_power += (box_number + 1) * (lens_index + 1) * focal_length
		}
	}

	return focusing_power
}

func day15_part2() any {
	input, err := os.ReadFile("resources/15/input.txt")
	if err != nil {
		panic(err)
	}
	input = []byte(strings.Trim(string(input), "\n"))

	initialization_sequence := strings.Split(string(input), ",")

	boxes := [256][][2]string{}
	place_lenses(&boxes, initialization_sequence)

	return calculate_focusing_power(boxes)
}

func init() {
	RegisterPuzzle(15, 2, day15_part2)
}
