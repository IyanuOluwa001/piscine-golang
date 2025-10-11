package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	sign := 1
	start := 0

	// Handle optional sign at the beginning
	if s[0] == '+' {
		sign = 1
		start = 1
	} else if s[0] == '-' {
		sign = -1
		start = 1
	}

	// If after handling sign the string is empty, invalid input
	if start == len(s) {
		return 0
	}

	result := 0
	for i := start; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			return 0 // invalid character
		}
		digit := int(ch - '0')
		result = result*10 + digit
	}

	return sign * result
}
