package sudoku

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// CreateBoard = Get the sudoku puzzle from filename and input data into the board
func CreateBoard(filename string) (*Board, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		os.Exit(1)
	}
	defer file.Close()
	getSudokuPuzzle := bufio.NewScanner(file)
	board := new(Board)
	currentRow := 0

	for ; getSudokuPuzzle.Scan() && currentRow < SizeN; currentRow++ {
		//Remove all spaces of the current row
		rowTrimmed := strings.Split(strings.TrimSpace(getSudokuPuzzle.Text()), " ")

		fmt.Println(rowTrimmed)

		//Check if the input row fits into the board - size should be 9
		if len(rowTrimmed) != SizeN {
			return nil, fmt.Errorf("row %d needs to contain %d fields", currentRow, SizeN)
		}

		//Convert underscore and change them to 0 and change numbers from strings to int and insert them into the board
		//We are in the currentRow -- now time to browse the columns
		for currentColumn, val := range rowTrimmed {
			//For the column -> check if value is underscore and convert to zero
			cellValue, err := convertToDigit(val)
			if err != nil {
				return nil, fmt.Errorf("column %d needs to contain an underscore", currentColumn)
			}
			board.Cells[currentRow][currentColumn] = cellValue
		}
	}
	return board, nil
}

func convertToDigit(val string) (int, error) {
	//If val is underscore change to zero and return
	if val == "_" {
		return 0, nil
	}

	num, err := strconv.Atoi(val)
	if err != nil || num < 1 || num > SizeN {
		return 0, fmt.Errorf("only digits from 1 to %d and _ are allowed", SizeN)
	}

	return num, nil
}

// PrintBoard = takes the board and removes the brackets
func (b *Board) PrintBoard() string {
	sudokuBoard := ""

	for key, value := range b.Cells {
		str := fmt.Sprint(value)
		//Trim arrays brackets
		// sudokuBoard += str[1 : len(str)-1]
		sudokuBoard += str

		if key+1 < SizeN {
			sudokuBoard += "\n"
		}
	}
	// fmt.Println(sudokuBoard)
	return sudokuBoard
}
