package piscine

func Compact(ptr *[]string) int {
	s := *ptr
	count := 0

	// First count non-empty strings
	for _, v := range s {
		if v != "" {
			count++
		}
	}

	// Create a new slice with the correct size
	compacted := make([]string, count)
	index := 0

	// Copy non-empty strings into new slice
	for _, v := range s {
		if v != "" {
			compacted[index] = v
			index++
		}
	}

	// Update the original slice
	*ptr = compacted

	return count
}
