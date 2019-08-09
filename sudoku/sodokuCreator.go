package sudoku

func GenerateSudoku(hints int) Sudoku{
	sudoku := CreateEmptySudoku()

	// Generate Random Seed
	// Randomly fill Sudoku
	sudoku.randomisedBacktrackSolve()

	// Randomly remove values until total hints equals hints

	return sudoku
}

