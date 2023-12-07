package puzzles

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse_run(input string) Run {
	rows := strings.Split(input, "\n")

	time_parts := strings.Split(rows[0], " ")[1:]
	time_str := strings.Join(time_parts, "")

	distance_parts := strings.Split(rows[1], " ")[1:]
	distance_str := strings.Join(distance_parts, "")

	time, err := strconv.Atoi(time_str)
	if err != nil {
		panic(err)
	}

	distance, err := strconv.Atoi(distance_str)
	if err != nil {
		panic(err)
	}

	return Run{
		time:     time,
		distance: distance,
	}
}

func get_minimum_holding_time_for_record(run Run) (int, error) {
	for i := 0; i <= run.time; i++ {
		if distance_traveled, err := get_distance(i, run.time); err != nil {
			panic(err)
		} else if distance_traveled > run.distance {
			return i, nil
		}
	}

	return -1, fmt.Errorf("cannot find minimum holding time for record")
}

func get_maximum_holding_time_for_record(run Run) (int, error) {
	for i := run.time; i >= 0; i-- {
		if distance_traveled, err := get_distance(i, run.time); err != nil {
			panic(err)
		} else if distance_traveled > run.distance {
			return i, nil
		}
	}

	return -1, fmt.Errorf("cannot find maximum holding time for record")
}

func day6_part2() any {
	input, err := os.ReadFile("resources/6/input.txt")
	if err != nil {
		panic(err)
	}

	run := parse_run(string(input))
	minimum_holding_time_for_record, err := get_minimum_holding_time_for_record(run)
	if err != nil {
		panic(err)
	}
	maximum_holding_time_for_record, err := get_maximum_holding_time_for_record(run)
	if err != nil {
		panic(err)
	}
	number_of_possible_record_break := maximum_holding_time_for_record - minimum_holding_time_for_record + 1

	return number_of_possible_record_break
}

func init() {
	RegisterPuzzle(6, 2, day6_part2)
}
