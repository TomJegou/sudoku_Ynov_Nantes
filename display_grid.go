package sudoku

import "fmt"

func DisplayGrid(t [][]int) {
	//function that display the sudoku grid in a readable format
	for index_line := 0; index_line < len(t); index_line++ {
		fmt.Println(t[index_line])
	}
}
