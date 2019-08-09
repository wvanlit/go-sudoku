package main

import (
	"fmt"
	"sudoku/sudoku"
)

func main() {
	fmt.Println("Sudoku Version:", sudoku.Version)
	var grid = [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 3},
		{0, 0, 0, 0, 8, 7, 6, 4, 5},
		{0, 4, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 8, 0, 0, 9, 0, 0, 7},
		{0, 0, 7, 0, 0, 6, 5, 0, 2},
		{0, 0, 2, 0, 9, 0, 0, 0, 0},
		{0, 8, 0, 0, 0, 2, 0, 1, 6},
		{0, 3, 4, 0, 1, 0, 0, 7, 8},
	}
	var sud = sudoku.CreateSudoku(grid)
	println(sud.Solve())
	sud.PrintSudoku()
	//print(sud.FoundInRow(7,2))
}
