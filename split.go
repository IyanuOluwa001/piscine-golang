package piscine

func Split(s, sep string) []string {
	var result []string
	word := ""

	for i := 0; i < len(s); {
		// Check if the substring at position i starts with sep
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			// Add the current word to result and reset it
			result = append(result, word)
			word = ""
			i += len(sep) // Skip over the separator
		} else {
			word += string(s[i])
			i++
		}
	}
	// Add the last word
	result = append(result, word)
	return result
}
