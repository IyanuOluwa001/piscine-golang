package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		// Handle corner case: nbr == MinInt (to avoid overflow)
		if nbr == -nbr {
			// MinInt is tricky in Go, but since int is at least 32-bit, we can use uint trick:
			// Cast to uint and handle as positive
			printBase(uint(-nbr), base)
			return
		}
		nbr = -nbr
	}

	printBase(uint(nbr), base)
}

func printBase(n uint, base string) {
	baseLen := uint(len(base))
	if n >= baseLen {
		printBase(n/baseLen, base)
	}
	z01.PrintRune(rune(base[n%baseLen]))
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' {
			return false
		}
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
