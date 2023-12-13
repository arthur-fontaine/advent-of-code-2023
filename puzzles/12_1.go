package puzzles

import (
	"arthur-fontaine/advent-of-code-2023/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Spring [2]string

func parse_spring(spring string) Spring {
	spring_parts := strings.Split(spring, " ")
	return Spring{spring_parts[0], spring_parts[1]}
}

var cache_resolve_spring = map[string]int{}

func count_resolve_spring_ways(spring []rune, groups []int, spring_index int, group_index int, current_group_index int) int {
	// Mainly reimplementation of https://github.com/patrickarmengol/advent-of-code/blob/f357870c547c752696f8476455f0e0e19da6e490/2023/day12/solution.go#L13

	cache_key := fmt.Sprintf("%q %v %d %d %d", spring, groups, spring_index, group_index, current_group_index)

	if cached_result, ok := cache_resolve_spring[cache_key]; ok {
		return cached_result
	}

	if spring_index == len(spring) {
		// When we reach the end of the spring

		if group_index == len(groups) && current_group_index == 0 {
			// When we iterated over all groupes and there is no group currently being iterated
			return 1
		} else if group_index == len(groups)-1 && current_group_index == groups[group_index] {
			// When there is a group currently being iterated, but we iterated all this group
			return 1
		} else {
			return 0
		}
	}

	ways := 0
	next_char := spring[spring_index]

	spring_index++

	if next_char == '.' || next_char == '?' {
		if current_group_index == 0 {
			// There is no group currently being iterated, so we do not need to increment anythis
			ways += count_resolve_spring_ways(spring, groups, spring_index, group_index, 0)
		} else if current_group_index > 0 && group_index < len(groups) && groups[group_index] == current_group_index {
			// There is a group currently being iterated. As there is a dot and the group is completed, we need to stop the iteration, and start iterating the next group
			ways += count_resolve_spring_ways(spring, groups, spring_index, group_index+1, 0)
		}
	}

	if next_char == '#' || next_char == '?' {
		// We are currently iterating over a group
		ways += count_resolve_spring_ways(spring, groups, spring_index, group_index, current_group_index+1)
	}

	// Cache the result
	cache_resolve_spring[cache_key] = ways

	return ways
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
		groups, _ := utils.MapArray(strings.Split(spring[1], ","), func(n string) (int, error) { return strconv.Atoi(n) })
		s += count_resolve_spring_ways([]rune(spring[0]), groups, 0, 0, 0)
	}

	return s
}

func init() {
	RegisterPuzzle(12, 1, day12_part1)
}
