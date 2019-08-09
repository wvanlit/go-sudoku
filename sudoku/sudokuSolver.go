package sudoku

func (s sudoku) Solve() bool {
	if !s.hasIllegalValues() {
		s.solved = s.backtrackSolve()
		return s.solved
	}
	// If there are any illegal values it's automatically unsolved
	return false
}

func (s sudoku) getNextUnassignedValue() (int, int) {
	// Loop through grid
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if s.GetValueOnPosition(x, y) == 0 {
				return x, y
			}
		}
	}

	// No Unassigned Values
	return -1, -1
}

func (s sudoku) backtrackSolve() bool {
	// Check if done
	if s.Validate() {
		return true
	}

	x, y := s.getNextUnassignedValue()

	// Try every possible combination
	for value := 1; value <= 9; value++ {
		if s.isValid(value, x, y) {
			s.SetValueOnPosition(value, x, y)
			// Recursively call this function again to continue solving
			if s.backtrackSolve() {
				return true
			}
			// If the solve failed we reset the position
			s.SetValueOnPosition(0, x, y)
		}
	}
	// If there is no valid digit then an earlier digit is probably assigned on the wrong position
	// So we return false to continue backtracking
	return false
}

func (s sudoku) randomisedBacktrackSolve() bool {
	// Check if done
	if s.Validate() {
		return true
	}
	x, y := s.getNextUnassignedValue()
	// Try every possible combination
	digits := [9]int{1,2,3,4,5,6,7,8,9}
	random.Shuffle(9, func(i, j int) {
		digits[i], digits[j] = digits[j], digits[i]
	})
	for _,value := range digits {
		if s.isValid(value, x, y) {
			s.SetValueOnPosition(value, x, y)
			// Recursively call this function again to continue solving
			if s.backtrackSolve() {
				return true
			}
			// If the solve failed we reset the position
			s.SetValueOnPosition(0, x, y)
		}
	}
	// If there is no valid digit then an earlier digit is probably assigned on the wrong position
	// So we return false to continue backtracking
	return false
}
