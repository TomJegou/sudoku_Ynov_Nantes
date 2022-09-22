package sudoku

func ResolveGrid(grid [][][]int) [][][]int {
	// Function that resolve the sudoku's grid that give in argument
	/* The function fill the empty cell who have only 1 possible value. If there is no empty cell with only
	1 possible value, it take the cell who have the least possible value and try the first possible value. */
	var (
		minPossible int = 9
		minPos      []int
		found       bool = false
	)
	for line, tab := range grid {
		for column, val := range tab {
			if len(val) > 1 {
				if len(val) < minPossible {
					minPossible = len(val)
					minPos = []int{line, column}
					found = true
				}
			}
		}
	}
	if found {
		grid[minPos[0]][minPos[1]] = []int{grid[minPos[0]][minPos[1]][0]}
		// ReTest all possibilities for all empty cell
		return ResolveGrid(grid)
	} else {
		return grid
	}
}

func in2(n int, t []int) bool {
	for i := 0; i < len(t); i++ {
		if t[i] == n {
			return true
		}
	}
	return false
}
