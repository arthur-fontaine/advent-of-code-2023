package puzzles

import (
	"arthur-fontaine/advent-of-code-2023/utils"
	"os"
	"strconv"
	"strings"
)

type Spring struct {
	raw     []string
	by_size []int
}

func parse_spring(spring string) Spring {
	spring_parts := strings.Split(spring, " ")

	raw := strings.Split(spring_parts[0], "")
	by_size, _ := utils.MapArray(strings.Split(spring_parts[1], ","), func(v string) (int, error) {
		return strconv.Atoi(v)
	})

	return Spring{
		raw:     raw,
		by_size: by_size,
	}
}

func validate_spring(spring Spring) bool {
	// All raw character should be not empty
	if utils.ArrayIncludes(spring.raw, "") {
		return false
	}

	sizes := []int{}

	for _, character := range spring.raw {
		if character == "#" {
			if len(sizes) == 0 {
				sizes = append(sizes, 0)
			}

			sizes[len(sizes)-1]++
		} else if len(sizes) == 0 || sizes[len(sizes)-1] > 0 {
			sizes = append(sizes, 0)
		}
	}

	if sizes[len(sizes)-1] == 0 {
		sizes = sizes[:len(sizes)-1]
	}

	return utils.ArraysAreSame(sizes, spring.by_size)
}

func get_spring_possibilities(spring Spring) []string {
	possibilities := []string{}

	var g = func(s string) {}
	g = func(s string) {
		if strings.Count(s, "?") == 0 {
			possibilities = append(possibilities, s)
		} else {
			g(strings.Replace(s, "?", ".", 1))
			g(strings.Replace(s, "?", "#", 1))
		}
	}

	g(strings.Join(spring.raw, ""))

	return possibilities
}

func brute_force_resolve_spring(spring Spring) []Spring {
	possibilities := get_spring_possibilities(spring)

	working_springs := []Spring{}

	for _, possibility := range possibilities {
		new_spring := spring
		new_spring.raw = strings.Split(possibility, "")

		if validate_spring(new_spring) {
			working_springs = append(working_springs, new_spring)
		}
	}

	return working_springs
}

func day12_part1() any {
	input, err := os.ReadFile("resources/12/input.txt")
	if err != nil {
		panic(err)
	}
	input = []byte(strings.Trim(string(input), "\n"))

	raw_springs := strings.Split(string(input), "\n")
	springs := []Spring{}
	for _, raw_spring := range raw_springs {
		springs = append(springs, parse_spring(raw_spring))
	}

	s := 0
	for _, spring := range springs {
		working_springs := brute_force_resolve_spring(spring)
		s += len(working_springs)
	}

	return s
}

func init() {
	RegisterPuzzle(12, 1, day12_part1)
}
