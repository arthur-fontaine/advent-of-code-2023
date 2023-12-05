package puzzles

import (
	"fmt"
	"os"
)

func (almanac_map AlmanacMap) get_source(destination int) int {
	for _, range_ := range almanac_map {
		if range_.destination <= destination && destination <= range_.destination+range_.length {
			return range_.source + (destination - range_.destination)
		}
	}

	return destination
}

func (almanac_map AlmanacMap) sort_by_destination() AlmanacMap {
	sorted_almanac_map := almanac_map

	for i := range sorted_almanac_map {
		if i == len(sorted_almanac_map)-1 {
			continue
		}

		if sorted_almanac_map[i].destination > sorted_almanac_map[i+1].destination {
			sorted_almanac_map[i], sorted_almanac_map[i+1] = sorted_almanac_map[i+1], sorted_almanac_map[i]
		}
	}

	return sorted_almanac_map
}

func get_seed_from_location(almanac Almanac, location int) (int, error) {
	humidity := almanac.humidity_to_location.get_source(location)
	temperature := almanac.temperature_to_humidity.get_source(humidity)
	light := almanac.light_to_temperature.get_source(temperature)
	water := almanac.water_to_light.get_source(light)
	fertilizer := almanac.fetilizer_to_water.get_source(water)
	soil := almanac.soil_to_fertilizer.get_source(fertilizer)
	seed := almanac.seed_to_soil.get_source(soil)

	for _, seed_range := range almanac.seed_ranges {
		if seed_range[0] <= seed && seed <= seed_range[1] {
			return seed, nil
		}
	}

	return seed, fmt.Errorf("the seed does not exist")
}

func day5_part2() any {
	input, err := os.ReadFile("resources/5/input.txt")
	if err != nil {
		panic(err)
	}

	almanac := parse_almanac(string(input))

	for _, humidity_to_location := range almanac.humidity_to_location.sort_by_destination() {
		for location := 0; location < humidity_to_location.destination+humidity_to_location.length; location++ {
			if _, err := get_seed_from_location(almanac, location); err == nil {
				return location
			}
		}
	}

	return nil
}

func init() {
	RegisterPuzzle(5, 2, day5_part2)
}
