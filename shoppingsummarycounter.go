package piscine

import "github.com/01-edu/z01"

func ShoppingSummaryCounter(str string) {
	summary := make(map[string]int)
	order := []string{}
	word := ""

	for _, r := range str {
		if r == ' ' {
			if word != "" {
				if _, exists := summary[word]; !exists {
					order = append(order, word)
				}
				summary[word]++
				word = ""
			}
		} else {
			word += string(r)
		}
	}

	if word != "" {
		if _, exists := summary[word]; !exists {
			order = append(order, word)
		}
		summary[word]++
	}

	for _, item := range order {
		for _, r := range item {
			z01.PrintRune(r)
		}
		z01.PrintRune(' ')
		z01.PrintRune('=')
		z01.PrintRune('>')
		z01.PrintRune(' ')

		printNumber(summary[item])
		z01.PrintRune('$')
		z01.PrintRune('\n')
	}
}

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
		digits = append(digits, rune('0'+n%10))
		n /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
