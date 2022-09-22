package sudoku

import "strconv"

func compute_number_of_possibility(s []string, x int, y int) int {
	result := -9
	for i := 0; i < len(s[y]); i++ {
		if string(s[y][i]) != "." {
			result += 1
		}
	}
	for i := 0; i < len(s); i++ {
		if string(s[i][x]) != "." {
			result += 1
		}
	}
	return result
}

func PreprocessGrid(s []string) [][]int {
	//function that transform the grid given as argument
	//in order to make simpler the data process for the IA
	table := [][]int{}
	for y := 0; y < len(s); y++ {
		a := []int{}
		for x := 0; x < len(s[y]); x++ {
			if string(s[y][x]) == "." {
				a = append(a, compute_number_of_possibility(s, x, y))
			} else {
				val, _ := strconv.Atoi(string(s[y][x]))
				a = append(a, val)
			}
		}
		table = append(table, a)
	}
	return table
}
