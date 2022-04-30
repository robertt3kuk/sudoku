package sudoku

func ParseIt(args []string, s [9][9]int) [9][9]int {
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args); j++ {
			if []rune(args[i])[j] > 48 && []rune(args[i])[j] < 58 {
				s[i][j] = int(rune(args[i][j]) - 48)
			} else {
				s[i][j] = int(0)
			}
		}
	}
	return s
}
