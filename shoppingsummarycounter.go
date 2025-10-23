package piscine

import "github.com/01-edu/z01"

func ShoppingSummaryCounter(str string) {
	summary := make(map[string]int)
	order := []string{}
	word := ""

	// manually split string by spaces
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

	// handle last word
	if word != "" {
		if _, exists := summary[word]; !exists {
			order = append(order, word)
		}
		summary[word]++
	}

	// print results
	for _, item := range order {
		// print word
		for _, r := range item {
			z01.PrintRune(r)
		}
		z01.PrintRune(' ')
		z01.PrintRune('=')
		z01.PrintRune('>')
		z01.PrintRune(' ')

		// print count as digits
		count := summary[item]
		printNumber(count)

		// print dollar sign and newline
		z01.PrintRune('$')
		z01.PrintRune('\n')
	}
}

// helper to print integer using z01.PrintRune
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

	// reverse digits
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
