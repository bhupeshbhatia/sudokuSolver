package main

import (
	"fmt"
	"os"

	"github.com/bhupeshbhatia/sudokuSolver/sudoku"
)

func main() {
	fmt.Println("======= Sudoku Puzzle ========")

	board, err := sudoku.CreateBoard("puzzles/hardsudoku-3.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	if board.SolveSudoku() {
		fmt.Println("___________Solved_________")
		fmt.Println(board.PrintBoard())
	}
	// fmt.Println(board)
}
