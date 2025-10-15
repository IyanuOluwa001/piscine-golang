package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digitsCount := [10]int{}
	for n > 0 {
		digit := n % 10
		digitsCount[digit]++
		n /= 10
	}

	for digit := 0; digit <= 9; digit++ {
		for count := 0; count < digitsCount[digit]; count++ {
			z01.PrintRune(rune(digit + '0'))

			/*
			   digitsCount[1] = 1
			   Inner loop runs once → prints '1'
			   '0' in ASCII = 48
			   1 + '0' = 49 → '1'
			   Output so far: 1
			*/
		}
	}
}

/*
| Step | n before | n % 10 | digitsCount updated | n after |
| ---- | -------- | ------ | ------------------- | ------- |
| 1    | 321      | 1      | digitsCount[1] = 1  | 32      |
| 2    | 32       | 2      | digitsCount[2] = 1  | 3       |
| 3    | 3        | 3      | digitsCount[3] = 1  | 0       |
*/
