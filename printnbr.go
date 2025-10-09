package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		// Handle min int correctly by splitting:
		if n == -2147483648 {
			// For 32-bit min int hardcode printing:
			// Print "2147483648" after '-'
			PrintNbr(214748364)
			PrintNbr(8)
			return
		}
		n = -n
	}
	if n > 9 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(rune('0' + n%10))
}
