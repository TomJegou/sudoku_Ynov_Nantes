package sudoku

func check_zone(x int, x_end int, y int, y_end int, t [][][]int) bool {
	//check the unicity in the zone 3*3
	a := []int{}
	for y := y; y <= y_end; y++ {
		for x := x; x <= x_end; x++ {
			a = append(a, t[y][x][0])
		}
	}
	for i := 0; i < len(a)-1; i++ {
		cursor_end := len(a) - 1
		for cursor_end := cursor_end; cursor_end > i; cursor_end-- {
			if a[i] == a[cursor_end] {
				return false
			}
		}
	}
	return true
}

func CheckSudoku(t [][][]int) bool {
	//Function that check if the grid is correctly resolved
	for y := 0; y < len(t); y++ {
		for x := 0; x < len(t[y]); x++ {
			cursor_end_line := len(t[y]) - 1
			cursor_end_column := len(t) - 1
			for x < cursor_end_line {
				if t[y][x][0] == t[y][cursor_end_line][0] {
					return false
				} else {
					cursor_end_line--
				}
			}
			for y < cursor_end_column {
				if t[y][x][0] == t[cursor_end_column][x][0] {
					return false
				} else {
					cursor_end_column--
				}
			}
			if x%3 == 0 && y%3 == 0 && x+2 < len(t[y]) && y+2 < len(t) {
				if !check_zone(x, x+2, y, y+2, t) {
					return false
				}
			}
		}
	}
	return true
}
