package piscine

func LoafOfBread(str string) string {
	// Check if string is less than 5 characters
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	// Process string without removing spaces, keeping original characters
	result := ""
	count := 0
	i := 0

	for i < len(str) {
		// Skip spaces
		for i < len(str) && str[i] == ' ' {
			i++
		}
		// If we reach the end while skipping spaces
		if i >= len(str) {
			break
		}
		// Start a new 5-character group
		group := ""
		groupLen := 0
		// Collect 5 non-space characters
		for i < len(str) && groupLen < 5 {
			if str[i] != ' ' {
				group += string(str[i])
				groupLen++
			}
			i++
		}
		// If we got a full group of 5 characters
		if groupLen == 5 {
			result += group
			count++
			// Skip the next non-space character if we haven't reached the end
			for i < len(str) && str[i] == ' ' {
				i++
			}
			// Skip one non-space character
			if i < len(str) {
				i++
			}
			// Add space after group if more characters remain
			if i < len(str) || (i == len(str) && count*6 < len(str)) {
				result += " "
			}
		}
	}

	// If no valid groups were formed
	if count == 0 {
		return "Invalid Output\n"
	}

	return result + "\n"
}
