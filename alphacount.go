package piscine

func AlphaCount(s string) int {
	iyanu := 0
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			iyanu++
		}
	}
	return iyanu
}
