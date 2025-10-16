package piscine

func ConcatParams(args []string) string {
	var result string

	for i, word := range args {
		result += word
		if i != len(args)-1 {
			result += "\n"
		}
	}

	return result
}
