package puzzles

import (
	"arthur-fontaine/advent-of-code-2023/utils"
	"os"
	"strings"
)

var play_cycle func(platform string) string

func define_play_cycle() {
	play_cycle = utils.Memoize(func(platform string) string {
		// Play north
		platform = up_all_rounded_rock(platform)

		// Rotate from north to west
		platform = utils.RotateString(platform)
		// Play west
		platform = up_all_rounded_rock(platform)

		// Rotate from west to south
		rows, _ := utils.MapArray(
			utils.ReverseArray(
				strings.Split(
					utils.RotateString(platform),
					"\n",
				),
			),
			func(s string) (string, error) {
				return utils.Reverse(s), nil
			},
		)
		platform = strings.Join(rows, "\n")
		// Play south
		platform = up_all_rounded_rock(platform)

		// Rotate from south to east
		platform = utils.RotateString(platform)
		// Play east
		platform = up_all_rounded_rock(platform)

		// Reset to north
		rows, _ = utils.MapArray(
			utils.ReverseArray(
				strings.Split(
					utils.RotateString(platform),
					"\n",
				),
			),
			func(s string) (string, error) {
				return utils.Reverse(s), nil
			},
		)
		platform = strings.Join(rows, "\n")

		return platform
	})
}

func day14_part2() any {
	input, err := os.ReadFile("resources/14/input.txt")
	if err != nil {
		panic(err)
	}
	input = []byte(strings.Trim(string(input), "\n"))

	define_get_platform_score()
	define_play_cycle()

	platform := string(input)
	for i := 1; i <= 1_000_000_000; i++ {
		platform = play_cycle(platform)
	}

	return get_platform_score(platform)
}

func init() {
	RegisterPuzzle(14, 2, day14_part2)
}
