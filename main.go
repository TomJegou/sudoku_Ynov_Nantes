package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	a := PreprocessGrid(args)
	BackTracking(a, 0)
	DisplayGrid(a)
}
