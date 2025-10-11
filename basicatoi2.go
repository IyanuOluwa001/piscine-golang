package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' { // If character is not a digit
			return 0 // Return 0 immediately
		}
		digit := int(ch - '0')
		result = result*10 + digit
	}
	return result
}
