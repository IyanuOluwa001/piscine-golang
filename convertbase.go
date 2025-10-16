package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Step 1: Convert nbr (string) in baseFrom → integer (base 10)
	num := 0
	baseLenFrom := len(baseFrom)

	for _, ch := range nbr {
		num = num*baseLenFrom + indexOf(ch, baseFrom)
	}

	// Step 2: Convert integer (base 10) → baseTo representation
	if num == 0 {
		return string(baseTo[0])
	}

	result := ""
	baseLenTo := len(baseTo)

	for num > 0 {
		remainder := num % baseLenTo
		result = string(baseTo[remainder]) + result
		num /= baseLenTo
	}

	return result
}

// Helper: find index of a rune in a string
func indexOf(ch rune, base string) int {
	for i, b := range base {
		if b == ch {
			return i
		}
	}
	return -1
}
