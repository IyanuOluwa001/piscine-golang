package main

import (
	"fmt"

	"piscine"
)

func main() {
	arg := 4
	fmt.Println(piscine.IterativeFactorial(arg))
	fmt.Println(piscine.IterativeFactorial1(arg))
}
