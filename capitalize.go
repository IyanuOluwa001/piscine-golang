package piscine

func Capitalize(s string) string {
	result := []rune{}
	inWord := false

	for _, r := range s {
		if isAlnum(r) {
			if !inWord {
				// Start of a new word: capitalize if letter
				if r >= 'a' && r <= 'z' {
					r = r - ('a' - 'A')
				}
				inWord = true
			} else {
				// Inside a word: lowercase if uppercase letter
				if r >= 'A' && r <= 'Z' {
					r = r + ('a' - 'A')
				}
			}
		} else {
			inWord = false
		}
		result = append(result, r)
	}

	return string(result)
}

func isAlnum(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9')
}
