package puzzles

import (
	"fmt"
	"os"
	"strings"
)

type TileWalker struct {
	type_          string
	row_index      int
	column_index   int
	pipe           Pipe
	outside_loop   bool
	in_main_loop   bool
	is_placeholder bool
}

type GridWalker [][]TileWalker

func (g GridWalker) to_string() string {
	string_grid := ""

	for _, row := range g {
		for _, tile := range row {
			if tile.type_ == "pipe" {
				string_grid += tile.pipe.symbol
			} else if tile.type_ == "ground" {
				string_grid += "."
			} else {
				panic("unknown tile")
			}
		}

		string_grid += "\n"
	}

	return strings.Trim(string_grid, "\n")
}

func parse_grid(string_grid string) (GridWalker, error) {
	row_strings := strings.Split(string_grid, "\n")

	grid := make([][]TileWalker, len(row_strings))

	for row_index, row_string := range row_strings {
		row := make([]TileWalker, len(row_string))

		for column_index, tile_rune := range row_string {
			tile_string := string(tile_rune)

			switch tile_string {
			case "|", "-", "L", "J", "7", "F":
				{
					pipe, err := find_pipe_by_symbol(tile_string)

					if err != nil {
						return GridWalker{}, err
					}

					row[column_index] = TileWalker{
						type_:        "pipe",
						pipe:         pipe,
						row_index:    row_index,
						column_index: column_index,
					}
				}
			case "/":
				{
					pipe, err := find_pipe_by_symbol("|")

					if err != nil {
						return GridWalker{}, err
					}

					row[column_index] = TileWalker{
						type_:          "pipe",
						pipe:           pipe,
						row_index:      row_index,
						column_index:   column_index,
						is_placeholder: true,
					}
				}
			case "_":
				{
					pipe, err := find_pipe_by_symbol("-")

					if err != nil {
						return GridWalker{}, err
					}

					row[column_index] = TileWalker{
						type_:          "pipe",
						pipe:           pipe,
						row_index:      row_index,
						column_index:   column_index,
						is_placeholder: true,
					}
				}
			case "S":
				{
					row[column_index] = TileWalker{
						type_:        "pipe",
						row_index:    row_index,
						column_index: column_index,
					}
				}
			case ".":
				{
					row[column_index] = TileWalker{
						type_:        "ground",
						row_index:    row_index,
						column_index: column_index,
					}
				}
			case "#":
				{
					row[column_index] = TileWalker{
						type_:          "ground",
						row_index:      row_index,
						column_index:   column_index,
						is_placeholder: true,
					}
				}
			default:
				{
					return GridWalker{}, fmt.Errorf("unknown tile %v", tile_string)
				}
			}
		}

		grid[row_index] = row
	}

	return grid, nil
}

func get_adjacent_positions(pipes []Pipe, tile_row_index int, tile_column_index int) struct {
	top    Pipe
	bottom Pipe
	left   Pipe
	right  Pipe
} {
	positions := struct {
		top    Pipe
		bottom Pipe
		left   Pipe
		right  Pipe
	}{}

	for _, pipe := range pipes {
		if pipe.row_index == tile_row_index-1 && pipe.column_index == tile_column_index {
			positions.top = pipe
		} else if pipe.row_index == tile_row_index+1 && pipe.column_index == tile_column_index {
			positions.bottom = pipe
		} else if pipe.row_index == tile_row_index && pipe.column_index == tile_column_index+1 {
			positions.right = pipe
		} else if pipe.row_index == tile_row_index && pipe.column_index == tile_column_index-1 {
			positions.left = pipe
		}
	}

	return positions
}

func get_s_replacement(string_grid string) string {
	row_strings := strings.Split(string_grid, "\n")

	for row_index, row_string := range row_strings {
		s_index := strings.Index(row_string, "S")

		if s_index != -1 {
			replacement := "."

			adjacent_pipes := get_adjacent_pipes(string_grid, row_index, s_index)
			adjacent_pipes_positions := get_adjacent_positions(adjacent_pipes, row_index, s_index)
			if adjacent_pipes_positions.top.symbol != "" {
				if adjacent_pipes_positions.bottom.symbol != "" {
					replacement = "|"
				} else if adjacent_pipes_positions.left.symbol != "" {
					replacement = "J"
				} else if adjacent_pipes_positions.right.symbol != "" {
					replacement = "L"
				}
			} else if adjacent_pipes_positions.bottom.symbol != "" {
				if adjacent_pipes_positions.left.symbol != "" {
					replacement = "7"
				} else if adjacent_pipes_positions.right.symbol != "" {
					replacement = "F"
				}
			} else if adjacent_pipes_positions.left.symbol != "" {
				if adjacent_pipes_positions.right.symbol != "" {
					replacement = "-"
				}
			}

			return replacement
		}
	}

	return "S"
}

func fill_main_loop(grid GridWalker, string_grid string, start_tile_row_index int, start_tile_column_index int) GridWalker {
	new_grid := grid

	new_grid[start_tile_row_index][start_tile_column_index].in_main_loop = true

	pipes_to_loop := [][]Pipe{get_adjacent_pipes(string_grid, start_tile_row_index, start_tile_column_index)}
	walked_indexes := map[int]map[int]bool{
		start_tile_row_index: {start_tile_column_index: true},
	}

	for i := 0; i < len(pipes_to_loop); i++ {
		next_pipes_to_loop := []Pipe{}
		for _, pipe := range pipes_to_loop[len(pipes_to_loop)-1] {
			new_grid[pipe.row_index][pipe.column_index].in_main_loop = true

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

	return new_grid
}

func replace_unused_pipes_by_ground(grid GridWalker, start_tile_row_index int, start_tile_column_index int) GridWalker {
	new_grid := grid
	string_grid := grid.to_string()

	new_grid = fill_main_loop(grid, string_grid, start_tile_row_index, start_tile_column_index)

	for row_index, row := range new_grid {
		for column_index, tile := range row {
			if !tile.in_main_loop && tile.type_ == "pipe" {
				new_grid[row_index][column_index] = TileWalker{
					type_:          "ground",
					row_index:      tile.row_index,
					column_index:   tile.column_index,
					is_placeholder: tile.is_placeholder,
				}
			}
		}
	}

	return new_grid
}

func get_placeholder_replacement(string_grid string, placeholder_row_index int, placeholder_column_index int) string {
	adjacent_pipes := get_adjacent_pipes(string_grid, placeholder_row_index, placeholder_column_index)

	// We only need to handle - and | pipes

	// This is a | pipe in cases:
	//  - there is (7 or F or |) at top and (J or L or |) at bottom

	// This is a - pipe in cases:
	//  - there is (7 or J or -) at right and (F or L or -) at left

	adjacent_pipes_positions := get_adjacent_positions(adjacent_pipes, placeholder_row_index, placeholder_column_index)

	if (adjacent_pipes_positions.top.symbol == "7" ||
		adjacent_pipes_positions.top.symbol == "F" ||
		adjacent_pipes_positions.top.symbol == "|") && (adjacent_pipes_positions.bottom.symbol == "J" ||
		adjacent_pipes_positions.bottom.symbol == "L" ||
		adjacent_pipes_positions.bottom.symbol == "|") {
		return "/"
	}

	if (adjacent_pipes_positions.left.symbol == "F" ||
		adjacent_pipes_positions.left.symbol == "L" ||
		adjacent_pipes_positions.left.symbol == "-") && (adjacent_pipes_positions.right.symbol == "7" ||
		adjacent_pipes_positions.right.symbol == "J" ||
		adjacent_pipes_positions.right.symbol == "-") {
		return "_"
	}

	return "#"
}

func add_placeholders_to_string_grid(string_grid string) string {
	new_grid := ""

	row_strings := strings.Split(string_grid, "\n")

	for i, row_string := range row_strings {
		new_row := strings.Join(strings.Split(row_string, ""), "#")
		new_grid += new_row

		if i < len(row_strings)-1 {
			new_grid += "\n" + strings.Repeat("#", len(new_row)) + "\n"
		}
	}

	return new_grid
}

func replace_placeholders_from_string_grid(string_grid string) string {
	new_grid := ""

	row_strings := strings.Split(string_grid, "\n")

	for row_index, row_string := range row_strings {
		for column_index, column_rune := range row_string {
			if string(column_rune) != "#" {
				new_grid += string(column_rune)
				continue
			}

			new_grid += get_placeholder_replacement(string_grid, row_index, column_index)
		}

		new_grid += "\n"
	}

	new_grid = strings.Trim(new_grid, "\n")

	return new_grid
}

func get_adjacent_tiles(grid GridWalker, tile_row_index int, tile_column_index int) []TileWalker {
	adjacent_tiles := []TileWalker{}

	for _, t := range []struct {
		condition        bool
		row_index        int
		column_index     int
		adjacent_symbols []string
	}{
		{condition: tile_row_index > 0, row_index: tile_row_index - 1, column_index: tile_column_index},
		{condition: tile_row_index < len(grid)-1, row_index: tile_row_index + 1, column_index: tile_column_index},
		{condition: tile_column_index > 0, row_index: tile_row_index, column_index: tile_column_index - 1},
		{condition: tile_column_index < len(grid[tile_row_index])-1, row_index: tile_row_index, column_index: tile_column_index + 1},
	} {
		if t.condition {
			adjacent_tiles = append(adjacent_tiles, grid[t.row_index][t.column_index])
		}
	}

	return adjacent_tiles
}

func walk_though_grid(grid *GridWalker, start_tile_row_index int, start_tile_column_index int) {
	tiles_to_loop := [][]TileWalker{get_adjacent_tiles(*grid, start_tile_row_index, start_tile_column_index)}
	walked_indexes := map[int]map[int]bool{
		start_tile_row_index: {start_tile_column_index: true},
	}

	for i := 0; i < len(tiles_to_loop); i++ {
		next_tiles_to_loop := map[TileWalker]bool{}
		for _, tile := range tiles_to_loop[len(tiles_to_loop)-1] {
			if walked_indexes[tile.row_index] == nil {
				walked_indexes[tile.row_index] = map[int]bool{}
			}
			walked_indexes[tile.row_index][tile.column_index] = true

			if tile.type_ != "ground" {
				continue
			}

			adjacent_grounds := []TileWalker{}
			for _, adjacent_tile := range get_adjacent_tiles(*grid, tile.row_index, tile.column_index) {
				if adjacent_tile.type_ == "ground" {
					adjacent_grounds = append(adjacent_grounds, adjacent_tile)
				}
			}

			for _, adjacent_ground := range adjacent_grounds {
				next_tiles_to_loop[adjacent_ground] = true
			}

			(*grid)[tile.row_index][tile.column_index].outside_loop = true
		}

		filtered_next_tiles_to_loop := []TileWalker{}
		for next_tile_to_loop := range next_tiles_to_loop {
			if !walked_indexes[next_tile_to_loop.row_index][next_tile_to_loop.column_index] {
				filtered_next_tiles_to_loop = append(filtered_next_tiles_to_loop, next_tile_to_loop)
			}
		}

		if len(filtered_next_tiles_to_loop) > 0 {
			tiles_to_loop = append(tiles_to_loop, filtered_next_tiles_to_loop)
		}
	}
}

func get_all_border_tiles(grid GridWalker) []TileWalker {
	border_tiles := []TileWalker{}

	border_tiles = append(border_tiles, grid[0]...)
	border_tiles = append(border_tiles, grid[len(grid)-1]...)

	for _, row := range grid[1 : len(grid)-1] {
		border_tiles = append(border_tiles, row[0])
		border_tiles = append(border_tiles, row[len(row)-1])
	}

	return border_tiles
}

func count_inside_tiles(grid GridWalker) int {
	inside_tiles_number := 0

	for _, row := range grid {
		for _, tile := range row {
			if tile.type_ == "ground" && !tile.outside_loop && !tile.is_placeholder {
				inside_tiles_number++
			}
		}
	}

	return inside_tiles_number
}

func day10_part2() any {
	input, err := os.ReadFile("resources/10/input.txt")
	if err != nil {
		panic(err)
	}
	input = []byte(strings.Trim(string(input), "\n"))

	s_replacement := get_s_replacement(string(input))

	string_grid := add_placeholders_to_string_grid(string(input))
	start_row_index, start_column_index, err := find_start_index(string_grid, "S")
	if err != nil {
		panic(err)
	}
	string_grid = strings.Replace(string_grid, "S", s_replacement, 1)
	string_grid = replace_placeholders_from_string_grid(string_grid)

	grid, err := parse_grid(string_grid)
	if err != nil {
		panic(err)
	}
	grid = replace_unused_pipes_by_ground(grid, start_row_index, start_column_index)

	border_tiles := get_all_border_tiles(grid)

	for _, border_tile := range border_tiles {
		if border_tile.type_ != "ground" {
			continue
		}

		walk_though_grid(&grid, border_tile.row_index, border_tile.column_index)
	}

	grid[start_row_index][start_column_index].pipe.symbol = "S"

	return count_inside_tiles(grid)
}

func init() {
	RegisterPuzzle(10, 2, day10_part2)
}
