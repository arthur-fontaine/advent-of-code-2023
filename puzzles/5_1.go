package puzzles

import (
	"os"
	"strconv"
	"strings"
)

type AlmanacMap []struct {
	source      int
	destination int
	length      int
}

func (almanac_map AlmanacMap) get_destination(source int) int {
	for _, range_ := range almanac_map {
		if range_.source <= source && source <= range_.source+range_.length {
			return range_.destination + (source - range_.source)
		}
	}

	return source
}

type Almanac struct {
	seed_ranges             [][]int
	seed_to_soil            AlmanacMap
	soil_to_fertilizer      AlmanacMap
	fetilizer_to_water      AlmanacMap
	water_to_light          AlmanacMap
	light_to_temperature    AlmanacMap
	temperature_to_humidity AlmanacMap
	humidity_to_location    AlmanacMap
}

func get_almanac_part(almanac_input string, part_name string) string {
	almanac_input = " " + almanac_input
	number_list_str := strings.Split(strings.Split(almanac_input, part_name+":")[1], "\n\n")[0]
	number_list_str = strings.Trim(number_list_str, " ")

	return number_list_str
}

/*
Example:
convert_string_number_list("50 98 2") -> [50, 98, 2]
*/
func convert_string_number_list(string_number_list_str string) []int {
	string_number_list := strings.Split(string_number_list_str, " ")
	number_list := []int{}

	for _, string_number := range string_number_list {
		if len(string_number) == 0 {
			continue
		}

		number, err := strconv.Atoi(string_number)
		if err != nil {
			panic(err)
		}
		number_list = append(number_list, number)
	}

	return number_list
}

func parse_almanac_map(string_number_lists_str string) AlmanacMap {
	almanac_map := AlmanacMap{}

	string_number_lists_str = strings.Trim(strings.Trim(string_number_lists_str, " "), "\n")

	string_number_lists := strings.Split(string_number_lists_str, "\n")

	for _, string_number_list_str := range string_number_lists {
		number_list := convert_string_number_list(string_number_list_str)

		almanac_map = append(almanac_map, struct {
			source      int
			destination int
			length      int
		}{
			destination: number_list[0],
			source:      number_list[1],
			length:      number_list[2],
		})
	}

	return almanac_map
}

func parse_almanac(input string) Almanac {
	seed_flat_ranges := convert_string_number_list(get_almanac_part(input, "seeds"))
	seed_ranges := [][]int{}
	for i := 0; i < len(seed_flat_ranges); i += 2 {
		seed_ranges = append(seed_ranges, []int{seed_flat_ranges[i], seed_flat_ranges[i] + seed_flat_ranges[i+1]})
	}

	almanac := Almanac{
		seed_ranges:             seed_ranges,
		seed_to_soil:            parse_almanac_map(get_almanac_part(input, "seed-to-soil map")),
		soil_to_fertilizer:      parse_almanac_map(get_almanac_part(input, "soil-to-fertilizer map")),
		fetilizer_to_water:      parse_almanac_map(get_almanac_part(input, "fertilizer-to-water map")),
		water_to_light:          parse_almanac_map(get_almanac_part(input, "water-to-light map")),
		light_to_temperature:    parse_almanac_map(get_almanac_part(input, "light-to-temperature map")),
		temperature_to_humidity: parse_almanac_map(get_almanac_part(input, "temperature-to-humidity map")),
		humidity_to_location:    parse_almanac_map(get_almanac_part(input, "humidity-to-location map")),
	}

	return almanac
}

func get_location_from_seed(almanac Almanac, seed int) int {
	soil := almanac.seed_to_soil.get_destination(seed)
	fertilizer := almanac.soil_to_fertilizer.get_destination(soil)
	water := almanac.fetilizer_to_water.get_destination(fertilizer)
	light := almanac.water_to_light.get_destination(water)
	temperature := almanac.light_to_temperature.get_destination(light)
	humidity := almanac.temperature_to_humidity.get_destination(temperature)
	location := almanac.humidity_to_location.get_destination(humidity)
	return location
}

func day5_part1() any {
	input, err := os.ReadFile("resources/5/input.txt")
	if err != nil {
		panic(err)
	}

	almanac := parse_almanac(string(input))

	lowest_location := -1
	for _, seed_range := range almanac.seed_ranges {
		location := get_location_from_seed(almanac, seed_range[0])
		if lowest_location == -1 || location < lowest_location {
			lowest_location = location
		}
	}

	return lowest_location
}

func init() {
	RegisterPuzzle(5, 1, day5_part1)
}
