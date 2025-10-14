package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0 // empty substring is always found at start
	}

	lenS := len(s)
	lenFind := len(toFind)

	if lenFind > lenS {
		return -1
	}

	for i := 0; i <= lenS-lenFind; i++ {
		if s[i:i+lenFind] == toFind {
			return i
		}
	}

	return -1
}
