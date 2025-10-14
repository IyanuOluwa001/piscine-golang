package piscine

import "github.com/01-edu/z01"

func printsolution(position [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(position[i] + '1')) // print each queen row + '1'
	}
	z01.PrintRune('\n')
}

func isValid(position [8]int, col, row int) int {
	for prev := 0; prev < col; prev++ {
		if position[prev] == row || // same row
			position[prev]-prev == row-col || // same diagonal "\"
			position[prev]+prev == row+col { // same diagonal "/"
			return 0
		}
	}
	return 1
}

func placeQueen(col int, position [8]int) {
	if col == 8 {
		printsolution(position)
		return
	}
	for row := 0; row < 8; row++ {
		if isValid(position, col, row) == 1 {
			position[col] = row
			placeQueen(col+1, position)
		}
	}
}

func EightQueens() {
	var position [8]int
	placeQueen(0, position)
}
