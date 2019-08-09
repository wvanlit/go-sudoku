package main

import (
	"fmt"
	"sudoku/sudoku"
)

func main() {
	fmt.Println("Sudoku Version:", sudoku.Version)
	var sud = sudoku.GenerateSudoku(17)
	//println(sud.Solve())
	sud.PrintSudoku()
	//print(sud.FoundInRow(7,2))
}
