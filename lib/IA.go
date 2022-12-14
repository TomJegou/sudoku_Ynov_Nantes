package lib

func LineSudoku(ligne []int, n int) bool {
	//Check if the n parameter is already in the current line
	//Return true if it's not in the line and false if it's in
	for i := 0; i < len(ligne); i++ {
		if n == ligne[i] {
			return false
		}
	}
	return true
}

func ColonneSudoku(colonne [][]int, idColonne int, n int) bool {
	//Check if the n parameter is already in the current column
	//Return true if it's not in the line and false if it's in
	for i := 0; i < len(colonne); i++ {
		if colonne[i][idColonne] == n {
			return false
		}
	}
	return true
}

func SquareSudoku(tab [][]int, x int, y int, digit int) bool {
	// function who return true if digit in square n°square
	// initialization of varriable
	indexLign := 3 * (x / 3)
	indexCol := 3 * (y / 3)
	for i := indexLign; i < indexLign+3; i++ {
		for j := indexCol; j < indexCol+3; j++ {
			if tab[i][j] == digit {
				return false
			}
		}
	}
	return true
}

func BackTracking(grid [][]int, pos int) bool {
	//Recursive function that serve as an IA to complete the sudoku grid
	if pos == 9*9 { //condition to stop the recursivity
		return true
	}
	posX := pos / 9
	posY := pos % 9
	if grid[posX][posY] > 0 {
		return BackTracking(grid, pos+1)
	}
	for i := 1; i <= 9; i++ {
		if LineSudoku(grid[posX], i) && ColonneSudoku(grid, posY, i) && SquareSudoku(grid, posX, posY, i) {
			grid[posX][posY] = i
			if BackTracking(grid, pos+1) {
				return true
			}
		}
	}
	grid[posX][posY] = 0
	return false
}
