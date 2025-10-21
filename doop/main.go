package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	aStr := os.Args[1]
	op := os.Args[2]
	bStr := os.Args[3]

	a, err1 := strconv.ParseInt(aStr, 10, 64)
	b, err2 := strconv.ParseInt(bStr, 10, 64)

	if err1 != nil || err2 != nil {
		return
	}

	var result int64

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("No division by 0")
			return
		}
		result = a / b
	case "%":
		if b == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		result = a % b
	default:
		return
	}

	// Overflow check
	if result > math.MaxInt64 || result < math.MinInt64 {
		return
	}

	fmt.Println(result)
}
