package sudoku

import "time"

func GenerateSudoku(hints int) Sudoku {
	sudoku := CreateEmptySudoku()

	// Generate Random Seed
	random.Seed(time.Now().UnixNano())
	// Randomly fill Sudoku
	sudoku.randomisedBacktrackSolve()

	grid := sudoku.GetGrid()
	rowCount := len(grid)
	filledSquareCount := len(grid) * len(grid[0])

	// Randomly remove values until total hints equals hints
	for filledSquareCount > hints {
		x := random.Intn(rowCount)
		y := random.Intn(rowCount)

		val := sudoku.GetValueOnPosition(x, y)
		if val == 0 {
			continue
		}

		sudoku.DeleteValueOnPosition(x, y)
		testGrid := *grid
		testSudoku := CreateSudoku(testGrid)
		if testSudoku.Solve(){
			filledSquareCount --
		}else{
			sudoku.SetValueOnPosition(val, x, y)
		}
	}

	return sudoku
}
