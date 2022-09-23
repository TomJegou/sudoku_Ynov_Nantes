package lib

import "fmt"

func DisplayGrid(t [][]int) {
	//function that display the sudoku grid in a readable format
	print("\n\n")
	for index_line := 0; index_line < len(t); index_line++ {
		for index_column := 0; index_column < len(t[index_line]); index_column++ {
			fmt.Printf("%v", t[index_line][index_column])
			if index_column != 8 {
				if index_column%3 == 2 {
					fmt.Printf(" |")
				}
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
		if index_line%3 == 2 && index_line != 8 {
			fmt.Printf("―――――――――――――――――――――\n")
		}
	}
}
