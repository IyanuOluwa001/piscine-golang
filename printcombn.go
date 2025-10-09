package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	var comb []int
	printCombNHelper(n, 0, comb, true)
	z01.PrintRune('\n')
}

func printCombNHelper(n, start int, comb []int, first bool) {
	if len(comb) == n {
		if !first {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		for _, digit := range comb {
			z01.PrintRune(rune('0' + digit))
		}
		return
	}
	for i := start; i <= 10-(n-len(comb)); i++ {
		printCombNHelper(n, i+1, append(comb, i), first && len(comb) == 0)
		first = false
	}
}
