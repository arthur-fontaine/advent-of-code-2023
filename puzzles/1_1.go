package puzzles

import (
	"fmt"
)

func trebuchet() any {
	fmt.Println("Trebuchet")
	return 0
}

func init() {
	RegisterPuzzle(1, 1, trebuchet)
}
