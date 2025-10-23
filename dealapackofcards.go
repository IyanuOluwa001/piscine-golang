package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	// The deck is always 12 cards long
	cardsPerPlayer := 3
	player := 1
	index := 0

	for player <= 4 {
		fmt.Printf("Player %d: ", player)
		for j := 0; j < cardsPerPlayer; j++ {
			fmt.Printf("%d", deck[index])
			index++
			if j < cardsPerPlayer-1 {
				fmt.Printf(", ")
			}
		}
		z01.PrintRune('\n')
		player++
	}
}
