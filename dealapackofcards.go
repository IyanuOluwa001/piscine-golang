package piscine

import (
	"strconv"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	cardsPerPlayer := len(deck) / 4

	for i := 0; i < 4; i++ {
		player := "Player " + strconv.Itoa(i+1) + ": "
		for _, r := range player {
			z01.PrintRune(r)
		}

		start := i * cardsPerPlayer
		end := start + cardsPerPlayer

		for j := start; j < end; j++ {
			num := strconv.Itoa(deck[j])
			for _, r := range num {
				z01.PrintRune(r)
			}

			if j != end-1 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

/*
func DealAPackOfCards(deck []int) {

	cardsPerPlayer := len(deck) / 4

	for i := 0; i < 4; i++ {
		start := i * cardsPerPlayer
		end := start + cardsPerPlayer

		fmt.Printf("Player %d: ", i+1)
		for j := start; j < end; j++ {
			if j != end-1 {
				fmt.Printf("%d, ", deck[j])
			} else {
				fmt.Printf("%d", deck[j])
			}
		}
		z01.PrintRune()
	}
}
*/
