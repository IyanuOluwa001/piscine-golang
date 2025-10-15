package piscine

func IterativeFactorial1(nb int) int {
	if nb <= 0 {
		return 0
	}

	results := 1
	for i := 1; i < nb; i++ {
		results = results * i

		if results < 0 {
			return 0
		}
	}
	return results
}

// result is 1, i1, nb is 4
// r1 x i1 = 1
// r1 x i2 = r2
// r2 x i3 = r6
// r6 x i4 = r24

/*
package piscine

func FirstRune(s string) rune {
	for _, r := range s {
		return r // Return the first rune found
	}
	return 0 // Return 0 if string is empty (no rune)
}


- Go source text is Unicode characters encoded in UTF-8
*/
