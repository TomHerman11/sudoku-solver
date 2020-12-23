package main

import "fmt"

func main() {
	b := newSudokuBoardFromMatrix(getSudokuBoardExample(0))
	isValid, err := b.isSudokuBoardCellsValid()
	if !isValid {
		fmt.Println(err)
		return
	}

	isSolved := b.solve()
	if isSolved {
		fmt.Println("SOLVED:")
	} else {
		fmt.Println("FAILED. Could not solve the following board:")
	}
	b.printBoard()
}
