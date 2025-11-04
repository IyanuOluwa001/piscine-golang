package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ { // Loop for rows
		for j := 0; j < x; j++ { // Loop for columns
			if (i == 0 || i == y-1) && (j == 0 || j == x-1) {
				z01.PrintRune('o') // corners
			} else if i == 0 || i == y-1 {
				z01.PrintRune('-') // top/bottom edges
			} else if j == 0 || j == x-1 {
				z01.PrintRune('|') // side edges
			} else {
				z01.PrintRune(' ') // inside space
			}
		}
		z01.PrintRune('\n') // move to next line after each row
	}
}
