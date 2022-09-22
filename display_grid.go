package sudoku

import "fmt"

func DisplayGrid(t [][]int) {
	for index_line := 0; index_line < len(t); index_line++ {
		fmt.Println(t[index_line])
	}
}
