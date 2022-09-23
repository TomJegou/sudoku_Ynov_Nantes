package main

func check_zone(x int, x_end int, y int, y_end int, t [][]int) bool {
	//check the values's unicity in the 3*3 areas
	a := []int{}
	for y := y; y <= y_end; y++ {
		for x := x; x <= x_end; x++ {
			a = append(a, t[y][x])
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

func CheckSudoku(t [][]int) bool {
	//Function that check if the grid is correctly resolved
	for y := 0; y < len(t); y++ {
		for x := 0; x < len(t[y]); x++ {
			cursor_end_line := len(t[y]) - 1
			cursor_end_column := len(t) - 1
			for x < cursor_end_line {
				if t[y][x] == t[y][cursor_end_line] {
					return false
				} else {
					cursor_end_line--
				}
			}
			for y < cursor_end_column {
				if t[y][x] == t[cursor_end_column][x] {
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
