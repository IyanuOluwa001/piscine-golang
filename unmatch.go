package piscine

func Unmatch(a []int) int {
	counts := make(map[int]int)

	// Step 1: Count how many times each number appears
	for _, num := range a {
		counts[num]++
	}

	// Step 2: Return the first number in the original slice that appears an odd number of times
	for _, num := range a {
		if counts[num]%2 != 0 {
			return num
		}
	}

	// Step 3: If every number has a pair, return -1
	return -1
}
