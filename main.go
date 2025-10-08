package main

import "github.com/01-edu/z01"

/*(
	"fmt"
)*/

func main() {
	for x := 'a'; x <= 'z'; x++ {
		z01.PrintRune(x)
	}

	z01.PrintRune('\n')
}

/*
func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, r := range alphabet {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
*/
/*
func main() {
	fmt.Println("abcdefghijklmnopqrstuvwxyz")
}
*/
