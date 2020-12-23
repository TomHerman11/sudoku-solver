package main

import (
	"fmt"
	"math"
	"strings"
)

// DefaultBoardLength - yeah.
const DefaultBoardLength = 9

// SudokuBoard - A Sudoku board
type SudokuBoard struct {
	board  [][]int
	length int
}

func newEmptySudokuBoard() *SudokuBoard {
	if !isLengthValidForSudokuBoard(DefaultBoardLength) {
		panic(fmt.Errorf("Length: '%v' is not a square of an int, therefore it is not a valid length for a Sudoku board", DefaultBoardLength))
	}

	sudokuBoard := new(SudokuBoard)
	board := make([][]int, DefaultBoardLength)
	for i := range board {
		board[i] = make([]int, DefaultBoardLength)
	}

	sudokuBoard.board = board
	sudokuBoard.length = DefaultBoardLength
	return sudokuBoard
}

func newSudokuBoardFromMatrix(mat [][]int) *SudokuBoard {
	if !isMatrixDimensionsValidForSudokuBoard(mat) {
		panic(fmt.Errorf("Matrix '%v' cannot be used as a Sudoku board", mat))
	}

	sudokuBoard := new(SudokuBoard)

	// copy mat into board:
	board := make([][]int, len(mat))
	for i := range board {
		board[i] = make([]int, len(mat[i]))
		copy(board[i], mat[i])
	}

	sudokuBoard.board = board
	sudokuBoard.length = len(mat)
	return sudokuBoard
}

func (b *SudokuBoard) getBoxLength() int {
	return int(math.Sqrt(float64(b.length)))
}

// returns box origin row, column
func (b *SudokuBoard) getBoxOriginIndices(rowIndex int, columnIndex int) (int, int) {
	var boxLength int = b.getBoxLength()
	return rowIndex - (rowIndex % boxLength), columnIndex - (columnIndex % boxLength)
}

func (b *SudokuBoard) getAllowedDigits() []int {
	digits := make([]int, b.length)
	for i := range digits {
		digits[i] = i + 1
	}
	return digits
}

func (b *SudokuBoard) printBoard() {
	valMaxLen := getNumberOfDigits(len(b.board))
	boxLength := b.getBoxLength()

	for i := range b.board {
		if (i%boxLength == 0) && i != 0 {
			fmt.Println()
		}
		for boxNum := 0; boxNum < boxLength; boxNum++ {
			for j := boxNum * boxLength; j < (boxNum+1)*boxLength; j++ {
				prefixSpaces := strings.Repeat(" ", valMaxLen-getNumberOfDigits(b.board[i][j]))
				fmt.Printf("%v%v ", prefixSpaces, b.board[i][j])
			}
			fmt.Print("  ")
		}
		fmt.Println()
	}
}
