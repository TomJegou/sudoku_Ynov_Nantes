package sudoku

func ResolveGrid(grid [][]int) [][]int {
	// Function that resolve the sudoku's grid that give in argument
	/* The function fill the empty cell who have only 1 possible value. If there is no empty cell with only
	1 possible value, it take the cell who have the least possible value and try the first possible value. */
	var hardlogic bool = false
	for i := 0; i < 2; i++ {
		for line, tab := range grid {
			for column, val := range tab {
				if len(val) == 1 {
					// Add the only possible value
					grid[line][column] = val[0]
				} else if len(val) > 0 && hardlogic {

				}
			}
		}
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
