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
	count := 0
	skip := false

	for _, r := range str {
		if r == ' ' {
			continue
		}
		if skip {
			skip = false
			continue
		}

		result += string(r)
		count++

		if count == 5 {
			result += " "
			count = 0
			skip = true
		}
	}

	// remove any trailing space
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	result += "\n"
	return result
}
