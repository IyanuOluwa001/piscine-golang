package piscine

func LoafOfBread(str string) string {
	length := 0
	for _, r := range str {
		if r != ' ' {
			length++
		}
	}
	if length == 0 {
		return "\n"
	}

	if length < 5 {
		return "Invalid Output\n"
	}

	result := ""
	word := ""
	count := 0
	skipNext := false

	for _, r := range str {
		if skipNext {
			skipNext = false
			continue
		}
		if r == ' ' {
			continue
		}

		word += string(r)
		count++

		if count == 5 {
			if result != "" {
				result += " "
			}
			result += word
			word = ""
			count = 0
			skipNext = true
		}
	}

	if word != "" {
		if result != "" {
			result += " "
		}
		result += word
	}

	return result + "\n"
}
