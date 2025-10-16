package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder (n int) {
	//check if passed argument n = 0 and return 0
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	//create a digitsCount Array slot 
	digitsCount := [10]int{}
	//check how many times each digit in n happened
	for n>0 {
		digit := n%10
		digitsCount[digit]++
		//increment digit
		n = n/10
	}

	for digit := 0; digit <= 9; digit ++{
		for count := 0; count < digitsCount[digit]; count ++{
			z01.PrintRune(rune(digit + '0'))
		}
	}
	
}