package main

import (
	"sudoku/sudoku"
)

func main() {
	var sud = sudoku.GenerateSudoku(20)
	sud.PrintSudoku()
	sud.Solve()
	sud.PrintSudoku()
}
