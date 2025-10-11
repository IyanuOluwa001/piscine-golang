package piscine

func BasicAtoi(s string) int {
	result := 0            // Initialize result variable
	for _, ch := range s { // Iterate over each character in the string
		digit := int(ch - '0')     // Convert rune character to integer digit
		result = result*10 + digit // Shift previous digits and add current digit
	}
	return result // Return the final integer value
}
