package main

// difficulty - 0 = easy, 1 = medium, 2 = hard, ...
func getSudokuBoardExample(difficulty int) [][]int {
	m := make(map[int][][]int)
	m[0] = [][]int{
		{0, 0, 3, 5, 0, 9, 4, 0, 0},
		{0, 1, 8, 0, 7, 0, 9, 5, 0},
		{5, 0, 0, 0, 6, 0, 0, 0, 8},

		{0, 4, 0, 1, 0, 7, 0, 2, 0},
		{0, 3, 9, 0, 0, 0, 8, 0, 0},
		{0, 6, 0, 8, 0, 3, 0, 9, 0},

		{9, 0, 0, 0, 0, 0, 0, 0, 6},
		{0, 7, 6, 0, 4, 0, 1, 8, 0},
		{0, 0, 2, 6, 0, 8, 3, 0, 0},
	}

	m[1] = [][]int{
		{0, 4, 0, 5, 0, 0, 0, 0, 0},
		{5, 0, 0, 0, 0, 0, 0, 1, 6},
		{0, 8, 7, 0, 0, 1, 0, 0, 0},

		{3, 0, 6, 4, 9, 0, 0, 2, 8},
		{0, 9, 0, 1, 0, 2, 0, 0, 0},
		{1, 5, 0, 0, 3, 8, 0, 0, 9},

		{0, 0, 0, 6, 0, 0, 5, 4, 0},
		{4, 1, 0, 0, 0, 0, 0, 0, 7},
		{0, 0, 5, 0, 0, 7, 0, 0, 0},
	}

	m[2] = [][]int{
		{5, 0, 0, 0, 0, 0, 2, 0, 7},
		{0, 0, 0, 0, 3, 0, 0, 1, 0},
		{2, 0, 8, 0, 0, 7, 0, 9, 0},

		{6, 0, 1, 0, 9, 0, 0, 0, 0},
		{7, 3, 0, 0, 0, 0, 0, 8, 0},
		{0, 0, 0, 0, 2, 0, 1, 0, 6},

		{0, 9, 0, 3, 0, 0, 5, 0, 4},
		{0, 2, 0, 4, 1, 0, 0, 0, 0},
		{4, 0, 6, 0, 0, 0, 0, 0, 1},
	}

	return m[difficulty]
}