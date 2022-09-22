package sudoku

import "strconv"

func Preprocess_grid(s []string) []int {
	table := []int{}
	for y := 0; y < len(s); y++ {
		a := []int{}
		for x := 0; x < len(s[y]); x++ {
			if string(s[y][x]) == "." {
				a = append(a, -1)
			} else {
				val, _ := strconv.Atoi(string(s[y][x]))
				a = append(a, val)
			}
		}
		table = append(table, a...)
	}
	return table
}
