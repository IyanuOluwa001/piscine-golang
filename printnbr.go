package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		// Special handling for min int (e.g. -2147483648)
		if n == -2147483648 {
			z01.PrintRune('2')
			n = -147483648
		} else {
			n = -n
		}
	}

	digits := []rune{}
	for n > 0 {
		digit := rune(n%10 + '0')
		digits = append([]rune{digit}, digits...)
		n /= 10
	}

	for _, d := range digits {
		z01.PrintRune(d)
	}
}
