package sudoku

import "sudoku"

func BackTracking(grid [][]int, pos int) bool {
	if pos == 9*9 {
		return true
	}
	posX := pos / 9
	posY := pos % 9
	if grid[posX][posY] > 0 {
		return BackTracking(grid, pos+1)
	}
	for i := 1; i <= 9; i++ {
		grid[posX][posY] = i
		if sudoku.Linesudoku(grid[posX]) && sudoku.Colonnesudoku(grid, posY) && BackTracking(grid, pos+1) {
			return true
		}
	}
	grid[posX][posY] = 0
	return false
}

func in2(n int, t []int) bool {
	for i := 0; i < len(t); i++ {
		if t[i] == n {
			return true
		}
	}
	return false
}
