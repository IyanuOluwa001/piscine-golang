package piscine

func IsPrintable(s string) bool {
	for _, r := range s {
		// Printable ASCII characters range from 32 (space) to 126 (~)
		if r < 32 || r > 126 {
			return false
		}
	}
	return true
}
