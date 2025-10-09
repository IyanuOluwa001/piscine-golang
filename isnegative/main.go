package main

import "github.com/01-edu/z01"

func PrintIfNe(n int) {
	if n < 0 {
		z01.PrintRune('T')
	}
	if n == 0 {
		z01.PrintRune('F')
	}

	if n > 0 {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}

func main() {
	PrintIfNe(-5)
	PrintIfNe(0)
	PrintIfNe(7)
}
