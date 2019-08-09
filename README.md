# go-sudoku
[![Go Report Card](https://goreportcard.com/badge/github.com/wvanlit/go-sudoku)](https://goreportcard.com/report/github.com/wvanlit/go-sudoku)

A Sudoku Solver and Generator for the Go programming language

## Use
`Sudokugrid` is a [9][9]int array, which represents the Sudoku and it's values.

`createSudoku(g Sudokugrid)` returns a Sudoku struct that contains all the implemented methods.

`createEmptySudoku()` returns a Sudoku struct that contains all the implemented methods but has a Grid full of 0's.

`Generate Sudoku(hints int)` creates a new Sudoku struct with a randomised grid containing _hints_ amount of digits.

### Sudoku Methods
You have setters and getters for both the total grid `SetGrid(g Sudokugrid)` & `GetGrid()` 
and individual values in the grid `GetValueOnPosition(x int, y int)` & `SetValueOnPosition(value int, x int, y int)`.

`PrintSudoku()` prints the Sudoku to the console.

`Validate()` checks if the Sudoku is solved.

`hasIllegalValues()` checks if the current Sudoku doesn't have a value twice in the same cell, row or column.

`Solve()` solves the Sudoku using a backtracking algorithm.



