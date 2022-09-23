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
		grid = Reprossessing(grid)
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

func Reprossessing(grid [][][]int) [][][]int {
	for index_line := 0; index_line < len(grid); index_line++ {
		for index_column := 0; index_column < len(grid[index_line]); index_column++ {
			if len(grid[index_line][index_column]) > 1 {
				grid[index_line][index_column] = CheckPossibilities(grid, index_line, index_column)
			}
		}
	}
	return grid
}

func CheckPossibilities(grid [][][]int, line, column int) []int {
	listVals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for index_line := 0; index_line < len(grid); index_line++ {
		if index_line != line {
			if len(grid[index_line][column]) == 1 {
				if in2(grid[index_line][column][0], listVals) {
					PopList(listVals, grid[index_line][column][0])
				}
			}
		}
	}
	for index_column := 0; index_column < len(grid[0]); index_column++ {
		if index_column != column {
			if len(grid[line][index_column]) == 1 {
				if in2(grid[line][index_column][0], listVals) {
					PopList(listVals, grid[line][index_column][0])
				}
			}
		}
	}
	return listVals
}

func PopList(list []int, val int) []int {
	list2 := []int{}
	for _, val2 := range list {
		if val2 != val {
			list2 = append(list2, val2)
		}
	}
	return list2
}
