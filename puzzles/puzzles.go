package puzzles

type PuzzleIdentifier struct {
	Day  int
	Part int
}

var Puzzles = map[PuzzleIdentifier]func() any{}

func RegisterPuzzle(day int, part int, f func() any) {
	Puzzles[PuzzleIdentifier{day, part}] = f
}

func RunPuzzle(day int, part int) any {
	return Puzzles[PuzzleIdentifier{day, part}]()
}
