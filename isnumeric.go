package piscine

func IsNumeric(s string) bool {
	if s == "" {
		return false // empty string is not numeric
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
