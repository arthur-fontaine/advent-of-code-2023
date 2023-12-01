package main

import (
	"arthur-fontaine/advent-of-code-2023/puzzles"
	"fmt"
	"os"
	"strconv"
)

func main() {
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	part, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Running day %d, part %d\n", day, part)
	puzzle_result := puzzles.RunPuzzle(day, part)

	fmt.Printf("Result: %v\n", puzzle_result)
}
