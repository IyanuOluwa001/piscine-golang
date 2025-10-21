package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printNbr(n int) {
	zero := '0'
	if n >= 10 {
		printNbr(n / 10)
	}
	digit := zero
	mod := n % 10
	for i := 0; i < mod; i++ {
		digit++
	}
	z01.PrintRune(digit)
}

func main() {
	points := &point{}
	setPoint(points)

	printStr("x = ")
	printNbr(points.x)
	printStr(", y = ")
	printNbr(points.y)
	printStr("\n")
}
