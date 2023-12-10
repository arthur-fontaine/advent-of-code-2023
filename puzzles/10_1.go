package puzzles

import (
	"arthur-fontaine/advent-of-code-2023/utils"
	"fmt"
	"os"
	"strings"
)

type Pipe struct {
	symbol       string
	row_index    int
	column_index int
}

var top_adjacent_pipe_symbols = []string{"|", "F", "7"}
var bottom_adjacent_pipe_symbols = []string{"|", "L", "J"}
var left_adjacent_pipe_symbols = []string{"-", "L", "F"}
var right_adjacent_pipe_symbols = []string{"-", "7", "J"}

var available_pipes = []Pipe{
	{symbol: "|"},
	{symbol: "-"},
	{symbol: "L"},
	{symbol: "J"},
	{symbol: "7"},
	{symbol: "F"},
}

func find_pipe_by_symbol(symbol string) (Pipe, error) {
	for _, pipe := range available_pipes {
		if pipe.symbol == symbol {
			pipe_copy := pipe

			return pipe_copy, nil
		}
	}

	return Pipe{}, fmt.Errorf("cannot find a pipe with the symbol %v", symbol)
}

func get_adjacent_pipes(string_grid string, tile_row_index int, tile_column_index int) []Pipe {
	row_strings := strings.Split(string_grid, "\n")

	adjacent_pipes := []Pipe{}

	for _, t := range []struct {
		condition        bool
		row_index        int
		column_index     int
		adjacent_symbols []string
	}{
		{adjacent_symbols: top_adjacent_pipe_symbols, condition: tile_row_index > 0, row_index: tile_row_index - 1, column_index: tile_column_index},
		{adjacent_symbols: bottom_adjacent_pipe_symbols, condition: tile_row_index < len(row_strings)-1, row_index: tile_row_index + 1, column_index: tile_column_index},
		{adjacent_symbols: left_adjacent_pipe_symbols, condition: tile_column_index > 0, row_index: tile_row_index, column_index: tile_column_index - 1},
		{adjacent_symbols: right_adjacent_pipe_symbols, condition: tile_column_index < len(row_strings[tile_row_index])-1, row_index: tile_row_index, column_index: tile_column_index + 1},
	} {
		if t.condition {
			adjacent_tile_string := string(row_strings[t.row_index][t.column_index])
			adjacent_pipe, err := find_pipe_by_symbol(adjacent_tile_string)

			if err == nil {
				adjacent_pipe.row_index = t.row_index
				adjacent_pipe.column_index = t.column_index

				if utils.ArrayIncludes(t.adjacent_symbols, adjacent_pipe.symbol) {
					adjacent_pipes = append(adjacent_pipes, adjacent_pipe)
				}
			}
		}
	}

	return adjacent_pipes
}

func get_farthest_pipe(string_grid string, start_tile_row_index int, start_tile_column_index int) int {
	pipes_to_loop := [][]Pipe{get_adjacent_pipes(string_grid, start_tile_row_index, start_tile_column_index)}
	walked_indexes := map[int]map[int]bool{
		start_tile_row_index: {start_tile_column_index: true},
	}

	for i := 0; i < len(pipes_to_loop); i++ {
		next_pipes_to_loop := []Pipe{}
		for _, pipe := range pipes_to_loop[len(pipes_to_loop)-1] {
			next_pipes_to_loop = append(next_pipes_to_loop, get_adjacent_pipes(string_grid, pipe.row_index, pipe.column_index)...)

			if walked_indexes[pipe.row_index] == nil {
				walked_indexes[pipe.row_index] = map[int]bool{}
			}
			walked_indexes[pipe.row_index][pipe.column_index] = true
		}

		filtered_next_pipes_to_loop := []Pipe{}
		for _, next_pipe_to_loop := range next_pipes_to_loop {
			if !walked_indexes[next_pipe_to_loop.row_index][next_pipe_to_loop.column_index] {
				filtered_next_pipes_to_loop = append(filtered_next_pipes_to_loop, next_pipe_to_loop)
			}
		}

		if len(filtered_next_pipes_to_loop) > 0 {
			pipes_to_loop = append(pipes_to_loop, filtered_next_pipes_to_loop)
		}
	}

	return len(pipes_to_loop)
}

func find_start_index(string_grid string, symbol string) (int, int, error) {
	row_strings := strings.Split(string_grid, "\n")

	for row_index, row_string := range row_strings {
		for column_index, tile_rune := range row_string {
			tile_string := string(tile_rune)

			if tile_string == symbol {
				return row_index, column_index, nil
			}
		}
	}

	return -1, -1, fmt.Errorf("cannot find any symbol %v", symbol)
}

func day10_part1() any {
	input, err := os.ReadFile("resources/10/input.txt")
	if err != nil {
		panic(err)
	}
	input = []byte(strings.Trim(string(input), "\n"))

	row_index, column_index, err := find_start_index(string(input), "S")
	if err != nil {
		panic(err)
	}

	return get_farthest_pipe(string(input), row_index, column_index)
}

func init() {
	RegisterPuzzle(10, 1, day10_part1)
}
