package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // skip the program name

	// Loop backward from the last argument to the first
	for i := len(args) - 1; i >= 0; i-- {
		word := args[i]
		for _, r := range word {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
