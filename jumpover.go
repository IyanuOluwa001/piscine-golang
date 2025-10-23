package piscine

func JumpOver(str string) string {
	if str == "" {
		return "\n"
	}

	result := ""
	count := 0

	for _, r := range str {
		count++
		if count%3 == 0 {
			result = result + string(r)
		}
	}

	if result == "" {
		return "\n"
	}

	return result + "\n"
}
