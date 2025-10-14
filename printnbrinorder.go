package piscine

import "fmt"

func PrintNbrInOrder(n int) {
	if n == 0 {
		fmt.Print("0")
		return
	}

	// Count digits frequency
	digitsCount := [10]int{}
	for n > 0 {
		digit := n % 10
		digitsCount[digit]++
		n /= 10
	}

	// Print digits in ascending order
	for digit := 0; digit <= 9; digit++ {
		for count := 0; count < digitsCount[digit]; count++ {
			fmt.Print(digit)
		}
	}
}
