package sudoku

// SolveSudoku uses backtracking algorithm
func (b *Board) SolveSudoku() bool {
	row, column, emptyCell := b.isCellEmpty()

	if !emptyCell {
		return true
	}

	for numCandidate := 1; numCandidate <= SizeN; numCandidate++ {
		// fmt.Println(!b.isNumRepeated(row, column, numCandidate))
		if b.isNumValid(row, column, numCandidate) {
			b.Cells[row][column] = numCandidate

			if b.SolveSudoku() {
				return true
			}

			//Reset the cell
			b.Cells[row][column] = 0
		}
	}
	return false
}

//isCellEmpty checks if the cell has a value of 0. If so then it means the cell is empty and the value of that particular row and column are returned.
//Otherwise row, column are equal to 0 and value of false is returned.
func (b *Board) isCellEmpty() (int, int, bool) {
	for row := 0; row < SizeN; row++ {
		for column := 0; column < SizeN; column++ {
			//If the cells in the board have 0 as a value then it means the cell is empty
			if b.Cells[row][column] == 0 {
				return row, column, true
			}
		}
	}
	//Return cell not empty if a value other than 0 is found
	return 0, 0, false
}

// isNumRepeated checks whether a given number is already present in a particular row, column or a 3 x 3 section
func (b *Board) isNumValid(row, column, numCandidate int) bool {

	//Section 3 x 3 row and column
	//Dividing integer rounds down (ignores decimal places)
	secRow := row / 3 * 3
	secColumn := column / 3 * 3

	for numIterator := 0; numIterator < SizeN; numIterator++ {
		// fmt.Println(secRow+numIterator/3, secColumn+numIterator%3)
		if b.Cells[row][numIterator] == numCandidate || b.Cells[numIterator][column] == numCandidate || b.Cells[secRow+numIterator/3][secColumn+numIterator%3] == numCandidate {
			return false
		}
	}
	return true
}
