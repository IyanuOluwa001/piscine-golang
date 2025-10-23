package piscine

import "fmt"

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	menu := map[string]food{
		"burger":  {preptime: 15},
		"chips":   {preptime: 10},
		"nuggets": {preptime: 12},
	}

	item, exists := menu[order]
	if !exists {
		fmt.Println("404")
		return 0
	}
	return item.preptime
}
