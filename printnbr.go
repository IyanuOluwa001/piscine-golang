package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		// convert n to uint to safely handle MinInt
		u := uint(-n)
		printUint(u)
	} else {
		printUint(uint(n))
	}
}

func printUint(u uint) {
	if u > 9 {
		printUint(u / 10)
	}
	z01.PrintRune(rune('0' + u%10))
}
