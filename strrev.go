package piscine

func StrRev(s string) string {
	runes := []rune(s)         // Convert the input string to a slice of runes (Unicode code points)
	n := len(runes)            // Get the length of the rune slice
	for i := 0; i < n/2; i++ { // Loop from the start to the middle of the slice

		runes[i], runes[n-1-i] = runes[n-1-i], runes[i] // Swap the rune at position i with the rune at position (n-1-i)
	}
	return string(runes) // Convert the rune slice back to a string and return it
}
