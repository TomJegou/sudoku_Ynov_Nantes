package sudoku

func Testvalsudoku(ligne []int)bool { 
	lentab := len(ligne)
	n := []int{}
	for i := 0; i < lentab; i++ { 
		for j := 0; j < len(n); j++ {
			if ligne[i] == n[j] { // check if i and j are the same or not, if yes, it returns false
				return false
			} else {
				n = append(n, ligne[i]) //puts ligne[i] in n 
		}
	} 
	 
	}	
	return true 
}