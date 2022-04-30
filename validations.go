package sudoku

func Valid(r, c, number int, s [9][9]int) bool {
	// column
	for j := 0; j < len(s); j++ {
		if number == s[j][c] {
			return false
		}
	}
	// raw
	for i := 0; i < len(s); i++ {
		if number == s[r][i] {
			return false
		}
	}
	return true
}
