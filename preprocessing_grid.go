package sudoku

import "strconv"

func in(n string, t []string) bool {
	for i := 0; i < len(t); i++ {
		if string(t[i]) == n {
			return true
		}
	}
	return false
}

func compute_number_of_possibility(s []string, x int, y int) []int {
	a := []string{}
	result := []int{}
	for i := 0; i < len(s[y]); i++ {
		if string(s[y][i]) != "." {
			if !in(string(s[y][i]), a) {
				a = append(a, string(s[y][i]))
			} else {
				continue
			}
		}
	}
	for i := 0; i < len(s); i++ {
		if string(s[i][x]) != "." {
			if !in(string(s[i][x]), a) {
				a = append(a, string(s[i][x]))
			} else {
				continue
			}
		}
	}
	for i := 1; i < len(s); i++ {
		if !in(strconv.Itoa(i), a) {
			result = append(result, i)
		}
	}
	return result
}

func PreprocessGrid(s []string) [][][]int {
	//function that transform the grid given as argument
	//in order to make simpler the data process for the IA
	table := [][][]int{}
	for y := 0; y < len(s); y++ {
		a := [][]int{}
		for x := 0; x < len(s[y]); x++ {
			b := []int{}
			if string(s[y][x]) == "." {
				a = append(a, compute_number_of_possibility(s, x, y))
			} else {
				val, _ := strconv.Atoi(string(s[y][x]))
				b = append(b, val)
				a = append(a, b)
			}
		}
		table = append(table, a)
	}
	return table
}
