package lib

import "fmt"

func Validation(args []string) bool {
	var (
		arg string
		id  int
	)
	if len(args) != 9 {
		fmt.Printf("Unvalid lenght of arguments")
		return false
	}
	for id, arg = range args {
		if len(arg) != 9 {
			fmt.Printf("Unvalid lenght of the argument %v", id)
			return false
		}
		for index, val := range arg {
			if val < '0' || val > '9' {
				if val != '.' {
					fmt.Printf("Invalid Number at line: %v, index: %v", id, index)
					return false
				}
			}
		}
	}
	return true
}
