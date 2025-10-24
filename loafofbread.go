package piscine

import "github.com/01-edu/z01"

func LoafOfBread(str string) {
	// Count non-space characters
	countNonSpace := 0
	for _, r := range str {
		if r != ' ' {
			countNonSpace++
		}
	}

	// If less than 5 non-space characters → Invalid Output
	if countNonSpace < 5 {
		for _, r := range "Invalid Output\n" {
			z01.PrintRune(r)
		}
		return
	}

	count := 0    // counts characters in current "word"
	skip := false // whether to skip next character
	word := 0     // counts total chars to handle skips

	for _, r := range str {
		if r == ' ' {
			continue
		}
		if skip {
			skip = false
			continue
		}

		z01.PrintRune(r)
		count++
		word++

		if count == 5 {
			z01.PrintRune(' ')
			count = 0
			skip = true // skip next char
		}
	}
	z01.PrintRune('\n')
}
