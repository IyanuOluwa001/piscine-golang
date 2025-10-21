package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // skip the program name Hello World

	for _, word := range args {
		for _, r := range word {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

/*
package main

import (
	"os"

	"github.com/01-edu/z01"
	)

func main() {
	args := os.Args[1:]

	for _, word := range args {
		for _, r := range word {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
args:= os.Args[1]
args := os.Args
args := os.Args[1]

package piscine
("os"
"piscine")

func main(){
arg := os.Args[1]
for _, words := range args {
for _, r := range word {
z01.PrintRune (r)
}
z01.PrintRune ('\n')
}
}
*/
