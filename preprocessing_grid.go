package sudoku

import "strconv"

func in(n string, t []string) bool {
	//Function that check if the n parameter is in the table t
	//Return true if it's in it dans false if not
	for i := 0; i < len(t); i++ {
		if string(t[i]) == n {
			return true
		}
	}
	return false
}

func compute_all_possibility(s []string, x int, y int) []int {
	//Return in a slice of all possible value for this point
	value_already_used := []string{}
	result := []int{}
	//loop to check all the values already used in the current line
	for i := 0; i < len(s[y]); i++ {
		if string(s[y][i]) != "." {
			if !in(string(s[y][i]), value_already_used) {
				value_already_used = append(value_already_used, string(s[y][i]))
			} else {
				continue
			}
		}
	}
	//loop to check all the values already used in the current column
	for i := 0; i < len(s); i++ {
		if string(s[i][x]) != "." {
			if !in(string(s[i][x]), value_already_used) {
				value_already_used = append(value_already_used, string(s[i][x]))
			} else {
				continue
			}
		}
	}
	//loop to append in the result slice what value hasn't been used
	for i := 1; i < len(s); i++ {
		if !in(strconv.Itoa(i), value_already_used) {
			result = append(result, i)
		}
	}
	return result
}

func PreprocessGrid(s []string) [][][]int {
	//function that transform the grid given as argument
	//in order to make simpler the data process for the IA
	//for all the "." append a list of all possible value
	table := [][][]int{}
	for y := 0; y < len(s); y++ {
		a := [][]int{}
		for x := 0; x < len(s[y]); x++ {
			b := []int{}
			if string(s[y][x]) == "." {
				a = append(a, compute_all_possibility(s, x, y))
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
