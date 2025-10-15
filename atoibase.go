package piscine

func AtoiBase(s string, base string) int {
	// Check if base is valid
	if !isValidBase1(base) {
		return 0
	}

	baseLen := len(base)
	valueMap := make(map[rune]int)

	for i, char := range base {
		valueMap[char] = i
	}

	result := 0
	for _, char := range s {
		value, exists := valueMap[char]
		if !exists {
			return 0 // should never happen per instructions
		}
		result = result*baseLen + value
	}
	return result
}

// Helper to check base validity
func isValidBase1(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' {
			return false
		}
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}
