package piscine

import "github.com/01-edu/z01"

func StringToIntSlice(str string) []int {
	var result []int
	for _, r := range str {
		result = append(result, int(r))
	}
	for i, v := range result {
		printInt(v)
		if i < len(result)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
	return result
}

func printInt(n int) {
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
		digits = append([]rune{rune('0' + n%10)}, digits...)
		n /= 10
	}
	for _, d := range digits {
		z01.PrintRune(d)
	}
}
