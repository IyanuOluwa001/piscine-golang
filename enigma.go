package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	// Save all the values first
	valA := ***a
	valB := *b
	valC := *******c
	valD := ****d

	// Perform the swaps
	*******c = valA // a → c
	****d = valC    // c → d
	*b = valD       // d → b
	***a = valB     // b → a
}
