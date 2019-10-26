package sudoku

import (
	"fmt"
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().Unix()))

type SudokuGrid [9][9]int

// Private struct with all variables
type sudoku struct {
	grid   *SudokuGrid
	solved bool
	seed   int64
}

// Public Struct for creating objects
type Sudoku struct {
	sudoku
}

func CreateEmptySudoku() Sudoku {
	return Sudoku{
		sudoku{
			solved: false,
			grid:   &SudokuGrid{},
		},
	}
}
func CreateSudoku(g SudokuGrid) Sudoku {
	return Sudoku{
		sudoku{
			solved: false,
			grid:   &g,
		},
	}
}

func (s sudoku) GetGrid() *SudokuGrid {
	return s.grid
}
func (s sudoku) SetGrid(g SudokuGrid) {
	s.grid = &g
}

func (s sudoku) GetValueOnPosition(x int, y int) int {
	if y >= len(s.grid) || x >= len(s.grid){
		return 0
	}
	return s.grid[y][x]
}
func (s sudoku) SetValueOnPosition(value int, x int, y int) {
	if y < len(s.grid) || x < len(s.grid){
		s.grid[y][x] = value
	}
}
func (s sudoku) DeleteValueOnPosition(x int, y int){
	if y < len(s.grid) || x < len(s.grid){
		s.grid[y][x] = 0
	}
}

func (s sudoku) PrintSudoku() {
	fmt.Println("| - - - | - - - | - - - |")
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// Print a straight line every 3 digits
			if j%3 == 0 {
				fmt.Print("| ")
			}
			fmt.Print(s.grid[i][j], " ")
		}
		// Print a line every 3 rows
		if (i%3)-2 == 0 {
			fmt.Print("|\n| - - - | - - - | - - - ")
		}
		fmt.Println("|")
	}
}

/*
 *	Sudoku Validation
 */
func (s sudoku) Validate() bool {
	// Test if all cells only have 1-9 once
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if s.cellValue(i, j) != (1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9) {
				return false
			}
		}
	}
	// Test Every Column
	for column := 0; column < 9; column++ {
		if s.columnValue(column) != (1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9) {
			return false
		}
	}

	// Test Every Row
	for row := 0; row < 9; row++ {
		if s.rowValue(row) != (1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9) {
			return false
		}
	}

	return true
}
func (s sudoku) hasIllegalValues() bool {
	// TODO
	return false
}

func (s sudoku) cellValue(cellX int, cellY int) int {
	// Get Cell Starting Position
	xStart := cellX * 3
	yStart := cellY * 3

	var sum int = 1
	// Sum Cell Values
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum *= s.GetValueOnPosition(yStart+i, xStart+j)
		}
	}
	return sum
}
func (s sudoku) columnValue(columnIndex int) int {
	// Get Cell Starting Position
	var sum int = 1
	// Sum Cell Values
	for y := 0; y < 9; y++ {
		sum *= s.GetValueOnPosition(y, columnIndex)
	}
	return sum
}
func (s sudoku) rowValue(rowIndex int) int {
	// Get Cell Starting Position
	var sum int = 1
	// Sum Cell Values
	for x := 0; x < 9; x++ {
		sum *= s.GetValueOnPosition(rowIndex, x)
	}
	return sum
}

/*
 *	Value Checking
 */
func (s sudoku) foundInCell(value int, cellX int, cellY int) bool {
	// Get Cell Starting Position
	xStart := cellX * 3
	yStart := cellY * 3

	// Search Cell Values for value
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if value == s.GetValueOnPosition(xStart+i, yStart+j) {
				return true
			}
		}
	}
	return false
}
func (s sudoku) foundInColumn(value int, columnIndex int) bool {
	// Search Column for value
	for y := 0; y < 9; y++ {
		if value == s.GetValueOnPosition(columnIndex, y) {
			return true
		}
	}
	return false
}
func (s sudoku) foundInRow(value int, rowIndex int) bool {
	// Search Column for value
	for x := 0; x < 9; x++ {
		if value == s.GetValueOnPosition(x, rowIndex) {
			return true
		}
	}
	return false
}
func (s sudoku) isValid(value int, x int, y int) bool {
	return !s.foundInCell(value, x/3, y/3) &&
		!s.foundInColumn(value, x) &&
		!s.foundInRow(value, y)
}
