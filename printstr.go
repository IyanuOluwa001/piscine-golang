package piscine

import "fmt"

func PrintStr(s string) {
	for _, char := range s {
		fmt.Printf("%c\n", char)
	}
}
