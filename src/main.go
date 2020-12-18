package main

import "fmt"

func main() {
	b := newSudokuBoardFromMatrix(getSudokuBoardExample(2))
	// fmt.Println("RECEIVED:")
	// b.printBoard()
	b.solve()

	fmt.Println("\nSOLVED:")
	b.printBoard()
}
