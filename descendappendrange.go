package piscine

func DescendAppendRange(max, min int) []int {
	result := []int{} // ensures it's an empty (non-nil) slice

	if max <= min {
		return result
	}

	for i := max; i > min; i-- {
		result = append(result, i)
	}

	return result
}
