package main

import (
	"fmt"
	"os"
	"sudoku/lib"
)

func main() {
	args := os.Args[1:]
	if lib.Validation(args) {
		a := lib.PreprocessGrid(args)
		lib.BackTracking(a, 0)
		if lib.CheckSudoku(a) {
			lib.DisplayGrid(a)
		} else {
			fmt.Println("Error, unable to solve the grid")
		}
	}
}
