package puzzles

import (
	"arthur-fontaine/advent-of-code-2023/utils"
	"os"
	"strconv"
	"strings"
)

func parse_spring2(spring string) Spring {
	spring_parts := strings.Split(spring, " ")

	first_spring_part := spring_parts[0]
	spring_parts[0] = ""
	for i := 0; i < 5; i++ {
		spring_parts[0] += first_spring_part
		if i < 4 {
			spring_parts[0] += "?"
		}
	}

	spring_parts[1] = strings.Trim(strings.Repeat(spring_parts[1]+",", 5), ",")

	return Spring{spring_parts[0], spring_parts[1]}
}

func day12_part2() any {
	input, err := os.ReadFile("resources/12/input.txt")
	if err != nil {
		panic(err)
	}
	input = []byte(strings.Trim(string(input), "\n"))

	raw_springs := strings.Split(string(input), "\n")
	springs := []Spring{}
	for _, raw_spring := range raw_springs {
		springs = append(springs, parse_spring2(raw_spring))
	}

	s := 0
	for _, spring := range springs {
		groups, _ := utils.MapArray(strings.Split(spring[1], ","), func(n string) (int, error) { return strconv.Atoi(n) })
		s += count_resolve_spring_ways([]rune(spring[0]), groups, 0, 0, 0)
	}

	return s
}

func init() {
	RegisterPuzzle(12, 2, day12_part2)
}
