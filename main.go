package main

import (
	"os"
	"sudoku/lib"
)

func main() {
	args := os.Args[1:]
	a := lib.PreprocessGrid(args)
	lib.BackTracking(a, 0)
	lib.DisplayGrid(a)
}
