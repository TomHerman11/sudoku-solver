package main

import (
	"fmt"
)

func main() {
	b := newSudokuBoardFromMatrix(getSudokuBoardExample(0))
	isValid, err := b.isSudokuBoardCellsValid()
	if !isValid {
		fmt.Println("The Sudoku board is not valid:", err)
		return
	}

	// Print progress indicator while solving:
	shouldStop := make(chan bool)
	go printProgressIndicator(shouldStop)

	// Solve:
	isSolved := b.solve()
	shouldStop <- true

	if isSolved {
		fmt.Println("SOLVED:")
	} else {
		fmt.Println("FAILED. Could not solve the following board:")
	}
	b.printBoard()
}
