package puzzles

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func update_gear_ratios(mapping_gear_part_number *map[int][]int, input string, word_range []int) bool {
	rows := strings.Split(input, "\n")

	number_of_rows := len(rows)
	number_of_columns := len(input) / number_of_rows

	// Because \n are counted in the word_range, we need to normalize by substract by the row index
	shifting := word_range[0] / (number_of_columns + 1)
	real_word_range := []int{word_range[0] - shifting, word_range[1] - shifting}
	word_row_index := int(real_word_range[0] / number_of_columns)

	word_columns_range := []int{
		real_word_range[0] % number_of_columns,
		(real_word_range[0] % number_of_columns) + (real_word_range[1] - real_word_range[0]),
	}

	var min_row_index, max_row_index, min_column_index, max_column_index int

	if word_row_index == 0 {
		min_row_index = 0
	} else {
		min_row_index = word_row_index - 1
	}

	if word_row_index == number_of_rows-1 {
		max_row_index = number_of_rows - 1
	} else {
		max_row_index = word_row_index + 1
	}

	if word_columns_range[0] == 0 {
		min_column_index = 0
	} else {
		min_column_index = word_columns_range[0] - 1
	}

	if word_columns_range[1] == number_of_columns {
		max_column_index = number_of_columns - 1
	} else {
		max_column_index = word_columns_range[1]
	}

	characters := ""
	characters += rows[min_row_index][min_column_index : max_column_index+1]
	characters += rows[word_row_index][min_column_index : max_column_index+1]
	characters += rows[max_row_index][min_column_index : max_column_index+1]

	not_symbol_regex := regexp.MustCompile(`(?m)[^\d\.\s]`)

	gear_regex := regexp.MustCompile(`(?m)[\*]`)
	number, _ := strconv.Atoi(input[word_range[0]:word_range[1]])
	for _, row_index := range []int{min_row_index, word_row_index, max_row_index} {
		string_part := rows[row_index][min_column_index : max_column_index+1]
		if gear_match_ranges := gear_regex.FindAllStringIndex(string_part, -1); len(gear_match_ranges) > 0 {
			for _, gear_match_range := range gear_match_ranges {
				gear_id := row_index*number_of_columns + min_column_index + gear_match_range[0]
				(*mapping_gear_part_number)[gear_id] = append((*mapping_gear_part_number)[gear_id], number)
			}
		}
	}

	return len(not_symbol_regex.FindStringIndex(characters)) > 0
}

func day3_part2() any {
	input, err := os.ReadFile("resources/3/input.txt")
	if err != nil {
		panic(err)
	}

	number_indexes := get_number_indexes(string(input))
	mapping_gear_part_number := map[int][]int{}

	for _, number_index := range number_indexes {
		update_gear_ratios(&mapping_gear_part_number, string(input), number_index)
	}

	gear_ratio_sum := 0
	for _, part_numbers := range mapping_gear_part_number {
		if len(part_numbers) >= 2 {
			part_number_product := 1
			for _, part_number := range part_numbers {
				part_number_product *= part_number
			}
			gear_ratio_sum += part_number_product
		}
	}

	return gear_ratio_sum
}

func init() {
	RegisterPuzzle(3, 2, day3_part2)
}
