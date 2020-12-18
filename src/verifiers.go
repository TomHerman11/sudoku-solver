package main

import (
	"errors"
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

func (b *SudokuBoard) verifyRow(rowIndex int) (bool, error) {
	if rowIndex < 0 || rowIndex >= b.length {
		return false, errors.New("'rowIndex' is not valid")
	}

	m := make(map[int]bool)
	for i := 0; i < b.length; i++ {
		currVal := b.board[rowIndex][i]
		if currVal == 0 {
			// this value was never set / changed
			continue
		} else if currVal < 0 || currVal > b.length {
			return false, fmt.Errorf("The value '%v' at board[%v][%v] is not valid", currVal, rowIndex, i)
		} else if m[currVal] {
			return false, nil
		} else {
			m[currVal] = true
		}
	}
	return true, nil
}

func (b *SudokuBoard) verifyColumn(columnIndex int) (bool, error) {
	if columnIndex < 0 || columnIndex >= b.length {
		return false, errors.New("'columnIndex' is not valid")
	}

	m := make(map[int]bool)
	for i := 0; i < b.length; i++ {
		currVal := b.board[i][columnIndex]
		if currVal == 0 {
			// this value was never set / changed
			continue
		} else if currVal < 0 || currVal > b.length {
			return false, fmt.Errorf("The value '%v' at board[%v][%v] is not valid", currVal, i, columnIndex)
		} else if m[currVal] {
			return false, nil
		} else {
			m[currVal] = true
		}
	}
	return true, nil
}

func (b *SudokuBoard) verifyBox(rowIndex, columnIndex int) (bool, error) {
	if rowIndex < 0 || rowIndex >= b.length {
		return false, errors.New("'rowIndex' is not valid")
	} else if columnIndex < 0 || columnIndex >= b.length {
		return false, errors.New("'columnIndex' is not valid")
	}

	boxOriginRow, boxOriginColumn := b.getBoxOriginIndices(rowIndex, columnIndex)
	boxLength := b.getBoxLength()

	m := make(map[int]bool)
	for i := boxOriginRow; i < boxOriginRow+boxLength; i++ {
		for j := boxOriginColumn; j < boxOriginColumn+boxLength; j++ {
			currVal := b.board[i][j]
			if currVal == 0 {
				// this value was never set / changed
				continue
			} else if currVal < 0 || currVal > b.length {
				return false, fmt.Errorf("The value '%v' at board[%v][%v] is not valid", currVal, i, j)
			} else if m[currVal] {
				return false, nil
			} else {
				m[currVal] = true
			}
		}
	}
	return true, nil
}

func isLengthValidForSudokuBoard(l int) bool {
	lSqrt := math.Sqrt(float64(l))

	// check if the lSqrt is an integer:
	return lSqrt == float64(int(lSqrt))
}

func isMatrixDimensionsValidForSudokuBoard(mat [][]int) bool {
	l := len(mat)
	if !isLengthValidForSudokuBoard(l) {
		return false
	}

	for _, row := range mat {
		if len(row) != l {
			return false
		}
	}

	return true
}
