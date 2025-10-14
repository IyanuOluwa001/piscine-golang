package piscine

func FirstRune(s string) rune {
	for _, r := range s {
		return r // Return the first rune found
	}
	return 0 // Return 0 if string is empty (no rune)
}
