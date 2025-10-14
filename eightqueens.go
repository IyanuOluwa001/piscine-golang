package piscine

import "fmt"

// EightQueens prints all solutions to the 8-queens puzzle
func EightQueens() {
	positions := make([]int, 8) // positions[col] = row of queen in that column
	solve(positions, 0)
}

// solve tries to place queens starting from column 'col'
func solve(positions []int, col int) {
	if col == 8 {
		printSolution(positions)
		return
	}

	// Try placing queen in each row for this column
	for row := 1; row <= 8; row++ {
		if isSafe(positions, col, row) {
			positions[col] = row
			solve(positions, col+1)
		}
	}
}

// isSafe checks if a queen can be safely placed at (col,row)
func isSafe(positions []int, col, row int) bool {
	for c := 0; c < col; c++ {
		r := positions[c]
		if r == row ||                // same row
			r-c == row-col ||         // same diagonal "\"
			r+c == row+col {          // same diagonal "/"
			return false
		}
	}
	return true
}

// printSolution prints the positions as a string of digits
func printSolution(positions []int) {
	for _, row := range positions {
		fmt.Print(row)
	}
	fmt.Println()
}
