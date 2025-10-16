/*package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programPath := os.Args[0]
	name := ""

	// Find the last '/' in the path manually
	for i := len(programPath) - 1; i >= 0; i-- {
		if programPath[i] == '/' {
			name = programPath[i+1:]
			break
		}
	}

	// If there was no '/', the name is just the full path
	if name == "" {
		name = programPath
	}

	// Print each character
	for _, r := range name {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
*/

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	progPath := os.Args[0]

	lastSlash := -1
	for i, r := range progPath {
		if r == '/' {
			lastSlash = i
		}
	}

	progName := ""
	if lastSlash != -1 {
		progName = progPath[lastSlash+1:]
	} else {
		progName = progPath
	}

	for _, r := range progName {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
