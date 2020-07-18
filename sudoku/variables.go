package sudoku

// SizeN = Size of the board
const SizeN = 9

// Board =State of sudoku
type Board struct {
	Cells [SizeN][SizeN]int
}
