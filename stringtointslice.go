package piscine

import "github.com/01-edu/z01"

func StringToIntSlice(str string) {
	for _, r := range str {
		printNumber(int(r)) // print the integer value of each rune
		z01.PrintRune(' ')  // separate numbers with a space
	}
	z01.PrintRune('\n')
}

// helper to print an integer using z01.PrintRune
func printNumber(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	var digits []rune
	for n > 0 {
		digits = append(digits, rune('0'+(n%10)))
		n /= 10
	}

	// print digits in reverse
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
