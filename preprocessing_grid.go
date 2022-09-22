package sudoku

import "strconv"

func PreprocessGrid(s []string) [][]int {
	//function that transform the grid given as argument
	//in order to make simpler the data process for the IA
	table := [][]int{}
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
		table = append(table, a)
	}
	return table
}
