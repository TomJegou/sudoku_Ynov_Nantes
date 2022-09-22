package sudoku

func ResolveGrid(grid [][]int) [][]int {
	// Function that resolve the sudoku's grid that give in argument
	/* The function fill the empty cell who have only 1 possible value. If there is no empty cell with only
	1 possible value, it take the cell who have the least possible value and try the first possible value. */
	var (
		line   int
		column int
	)
	for line = 0; line < len(grid); line++ {
		for column = 0; column < len(grid[0]); column++ {
			if grid[line][column] == -1 {
				// Fill the cell with the only possible value
				// grid[line][column] = ** Function who retrun the first possible value **
			}
		}
	}
	return grid
}
