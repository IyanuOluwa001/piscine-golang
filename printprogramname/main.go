package main

import (
	"fmt"
	"os"
)

func main() {
	programPath := os.Args[0] // e.g. "./Nessy" or "/home/student/printprogramname"
	name := ""

	// Find the last '/' in the path manually
	for i := len(programPath) - 1; i >= 0; i-- {
		if programPath[i] == '/' {
			name = programPath[i+1:] // take everything after the last '/'
			break
		}
	}

	// If there was no '/', the name is just the full path
	if name == "" {
		name = programPath
	}

	fmt.Println(name)
}
