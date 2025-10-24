package piscine

func LoafOfBread(str string) string {
	// Count non-space characters
	countNonSpace := 0
	for _, r := range str {
		if r != ' ' {
			countNonSpace++
		}
	}

	// If less than 5 non-space characters → Invalid Output
	if countNonSpace < 5 {
		return "Invalid Output\n"
	}

	result := ""
	word := ""
	count := 0
	skipNext := false

	for _, r := range str {
		if r == ' ' {
			continue
		}
		if skipNext {
			skipNext = false
			continue
		}

		word += string(r)
		count++

		if count == 5 {
			// add group to result
			if len(result) > 0 {
				result += " "
			}
			result += word
			word = ""
			count = 0
			skipNext = true // skip next non-space character
		}
	}

	// add any remaining characters
	if len(word) > 0 {
		if len(result) > 0 {
			result += " "
		}
		result += word
	}

	result += "\n"
	return result
}
