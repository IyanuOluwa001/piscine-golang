package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	vowels := "aeiouAEIOU"
	allVowels := []rune{}

	// Step 1: Collect all vowels from all arguments
	for _, arg := range args {
		for _, r := range arg {
			if isVowel(r, vowels) {
				allVowels = append(allVowels, r)
			}
		}
	}

	// Step 2: Print each argument with vowels replaced (mirrored)
	vowelIndex := len(allVowels) - 1

	for i, arg := range args {
		for _, r := range arg {
			if isVowel(r, vowels) && vowelIndex >= 0 {
				z01.PrintRune(allVowels[vowelIndex])
				vowelIndex--
			} else {
				z01.PrintRune(r)
			}
		}
		if i < len(args)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func isVowel(r rune, vowels string) bool {
	for _, v := range vowels {
		if r == v {
			return true
		}
	}
	return false
}
