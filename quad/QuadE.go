package piscine

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			// Top-left corner
			if i == 0 && j == 0 {
				z01.PrintRune('A')
			// Top-right corner
			} else if i == 0 && j == x-1 {
				z01.PrintRune('C')
			// Bottom-left corner
			} else if i == y-1 && j == 0 {
				z01.PrintRune('C')
			// Bottom-right corner
			} else if i == y-1 && j == x-1 {
				z01.PrintRune('A')
			// Borders
			} else if i == 0 || i == y-1 || j == 0 || j == x-1 {
				z01.PrintRune('B')
			// Inside space
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
