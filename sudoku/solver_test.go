package sudoku

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkSolveSudoku(b *testing.B) {

	copy, err := CreateBoard("../puzzles/hardsudoku-2.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		// prevent the modification of the orignal board
		// copy := board.SolveSudoku
		b.StartTimer()
		copy.SolveSudoku()
	}
}
