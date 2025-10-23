package piscine

import "strings"

func ShoppingSummaryCounter(str string) map[string]int {
	// Step 1: Create a map to store item counts
	summary := make(map[string]int)

	// Step 2: Split the string into words (items)
	items := strings.Split(str, " ")

	// Step 3: Loop through each item and count
	for _, item := range items {
		// Increment the count for this item
		summary[item]++
	}

	// Step 4: Return the final summary map
	return summary
}
