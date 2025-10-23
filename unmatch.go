package piscine

func Unmatch(a []int) int {
	counts := make(map[int]int)

	// Count how many times each number appears
	for _, num := range a {
		counts[num]++
	}

	// Find the number that appears an odd number of times
	for key, val := range counts {
		if val%2 != 0 {
			return key
		}
	}

	// If all numbers have pairs, return -1
	return -1
}
