package sudoku

import "fmt"

func DisplayGrid(t [][][]int) {
	//function that display the sudoku grid in a readable format
	for index_line := 0; index_line < len(t); index_line++ {
		for index_column := 0; index_column < len(t[index_line]); index_column++ {
			fmt.Printf("%v", t[index_line][index_column][0])
			if index_column != 8 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
