package piscine

func StringToIntSlice(str string) []int {
	result := make([]int, len(str))
	for i, r := range str {
		result[i] = int(r)
	}
	return result
}
