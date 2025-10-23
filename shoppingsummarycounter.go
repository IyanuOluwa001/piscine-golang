package piscine

import "fmt"

func ShoppingSummaryCounter(str string) {
	summary := make(map[string]int)
	order := []string{} // to keep the insertion order
	word := ""

	// Split manually
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

	// handle last word (no trailing space)
	if word != "" {
		if _, exists := summary[word]; !exists {
			order = append(order, word)
		}
		summary[word]++
	}

	// Print each in order
	for _, item := range order {
		fmt.Printf("%s => %d$\n", item, summary[item])
	}
}
