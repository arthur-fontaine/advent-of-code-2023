package puzzles

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Run struct {
	time     int
	distance int
}

func parse_runs(input string) []Run {
	runs := []Run{}

	rows := strings.Split(input, "\n")

	times := strings.Split(rows[0], " ")[1:]
	distances := strings.Split(rows[1], " ")[1:]

	for _, time_str := range times {
		time_str := strings.Trim(time_str, " ")

		if len(time_str) == 0 {
			continue
		}

		if time, err := strconv.Atoi(time_str); err == nil {
			runs = append(runs, Run{
				time: time,
			})
		} else {
			panic(err)
		}
	}

	i := 0
	for _, distance_str := range distances {
		distance_str := strings.Trim(distance_str, " ")

		if len(distance_str) == 0 {
			continue
		}

		if distance, err := strconv.Atoi(distance_str); err == nil {
			runs[i].distance = distance
		} else {
			panic(err)
		}

		i++
	}

	return runs
}

func get_distance(load_milliseconds int, max_milliseconds int) (int, error) {
	if load_milliseconds > max_milliseconds {
		return 0, fmt.Errorf("load_milliseconds cannot be greater than max_milliseconds")
	}

	speed := load_milliseconds
	remaining_milliseconds := max_milliseconds - load_milliseconds

	// v = d/t
	// d = t*v

	return speed * remaining_milliseconds, nil
}

func get_number_of_possible_records(run Run) int {
	number_of_records := 0

	for i := 0; i <= run.time; i++ {
		if distance_traveled, err := get_distance(i, run.time); err != nil {
			panic(err)
		} else if distance_traveled > run.distance {
			number_of_records++
		}
	}

	return number_of_records
}

func day6_part1() any {
	input, err := os.ReadFile("resources/6/input.txt")
	if err != nil {
		panic(err)
	}

	runs := parse_runs(string(input))

	number_of_records_product := 1

	for _, run := range runs {
		number_of_records_product *= get_number_of_possible_records(run)
	}

	return number_of_records_product
}

func init() {
	RegisterPuzzle(6, 1, day6_part1)
}
