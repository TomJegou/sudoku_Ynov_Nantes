package lib

func Linesudoku(ligne []int, n int) bool {
	lentab := len(ligne)
	for i := 0; i < lentab; i++ {
		if n == ligne[i] {
			return false
		}
	}
	return true
}

func Colonnesudoku(colonne [][]int, idColonne int, n int) bool {
	lentab := len(colonne)
	for i := 0; i < lentab; i++ {
		if colonne[i][idColonne] == n {
			return false
		}
	}
	return true
}

func NotInSquare(tab [][]int, x int, y int, digit int) bool {
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
	if pos == 9*9 {
		return true
	}
	posX := pos / 9
	posY := pos % 9
	if grid[posX][posY] > 0 {
		return BackTracking(grid, pos+1)
	}
	for i := 1; i <= 9; i++ {
		if Linesudoku(grid[posX], i) && Colonnesudoku(grid, posY, i) && NotInSquare(grid, posX, posY, i) {
			grid[posX][posY] = i
			if BackTracking(grid, pos+1) {
				return true
			}
		}
	}
	grid[posX][posY] = 0
	return false
}
