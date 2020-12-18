package main

func (b *SudokuBoard) solve() {
	b.sudokuSolverRecGreedy(0, 0)
}

// Fill the board 'b' with a solution. Go over the indices
func (b *SudokuBoard) sudokuSolverRecGreedy(rowIndex, columnIndex int) bool {
	if rowIndex == -1 && columnIndex == -1 {
		// We reached the end of the board :D
		return true
	}

	nextRowIndex, nextColumnIndex := b.getNextCellIndicesToSolve(rowIndex, columnIndex)
	currVal := b.board[rowIndex][columnIndex]

	// If the current cell was not filled yet, try the different options:
	if currVal == 0 {
		options := b.getAllowedDigits()
		shuffleSliceInPlace(options)

		for _, v := range options {
			b.board[rowIndex][columnIndex] = v
			isRowValid, _ := b.verifyRow(rowIndex)
			isColumnValid, _ := b.verifyColumn(columnIndex)
			isBoxValid, _ := b.verifyBox(rowIndex, columnIndex)
			if isRowValid && isColumnValid && isBoxValid {
				if b.sudokuSolverRecGreedy(nextRowIndex, nextColumnIndex) {
					return true
				}
			}
		}
		b.board[rowIndex][columnIndex] = 0
		return false
	}

	// If this value was already given as part of the Sudoku, skip it:
	// fmt.Printf("currVal is '%v' at board[%v][%v] continuing...", currVal, rowIndex, columnIndex)
	// fmt.Println()
	return b.sudokuSolverRecGreedy(nextRowIndex, nextColumnIndex)
}

func (b *SudokuBoard) getNextCellIndicesToSolve(rowIndex, columnIndex int) (int, int) {
	nextRowIndex, nextColumnIndex := 0, 0
	if columnIndex < b.length-1 { // If this is NOT the end of the row:
		nextRowIndex = rowIndex
		nextColumnIndex = columnIndex + 1
	} else if columnIndex == b.length-1 { // If this is the end of the row:
		if rowIndex == b.length-1 {
			return -1, -1
		}
		nextRowIndex = rowIndex + 1
		nextColumnIndex = 0
	}
	return nextRowIndex, nextColumnIndex
}
