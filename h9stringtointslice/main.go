package main

import (
	"fmt"

	"piscine"
)

func main() {
	fmt.Println(piscine.StringToIntSlice("A quick brown fox jumps over the lazy dog"))
	fmt.Println(piscine.StringToIntSlice("Converted this string into an int"))
	fmt.Println(piscine.StringToIntSlice("hello THERE"))
}
