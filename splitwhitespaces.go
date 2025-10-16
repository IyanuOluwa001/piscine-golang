package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	var word string

	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(r)
		}
	}

	if word != "" {
		words = append(words, word)
	}

	return words
}
