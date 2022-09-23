package main

import (
	"fmt"
	"os"
	"sudoku/lib"
)

func main() {
	args := os.Args[1:]
	a := lib.PreprocessGrid(args)
	lib.BackTracking(a, 0)
	fmt.Println(lib.CheckSudoku(a))
	lib.DisplayGrid(a)
}
