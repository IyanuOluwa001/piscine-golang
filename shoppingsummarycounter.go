package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""

	for _, r := range str {
		if r == ' ' {
			if word != "" {
				summary[word]++
				word = ""
			}
		} else {
			word += string(r)
		}
	}

	// handle last word
	if word != "" {
		summary[word]++
	}

	return summary
}
