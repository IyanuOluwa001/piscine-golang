package piscine

func IsLower(s string) bool {
	if s == "" {
		return false // empty string is not lowercase
	}
	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}
