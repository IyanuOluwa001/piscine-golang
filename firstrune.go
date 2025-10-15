package piscine

func FirstRune(s string) rune {
	for _, r := range s {
		return r // Return the first rune found
	}
	return 0 // Return 0 if string is empty (no rune)
}

/*
Que 2: Write a function that returns the last rune of a string.

func LastRune(s string) rune {
	var last rune
	for _, r := range s {
		last = r
	}
	return last
}

Que 3:Write a function that returns the nth rune of a string.
If not possible, it returns 0.

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}

	i := 1
	for _, r := range s {
		if i == n {
			return r
		}
		i++
	}

	return 0 // n is larger than the number of runes
}

Que 4:
Write a function that behaves like the Compare function.

func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

Que 5:
Write a function that counts the letters of a string and returns the count.

func AlphaCount(s string) int {
	iyanu := 0
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			iyanu++
		}
	}
	return iyanu
}

Que 6:

*/
