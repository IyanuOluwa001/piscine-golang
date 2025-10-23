package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""

	for _, r := range str {
		if r == ' ' {
			// count current word, even if it's empty
			summary[word]++
			word = ""
		} else {
			word += string(r)
		}
	}

	// count last word
	summary[word]++

	return summary
}
