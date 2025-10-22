package main

import "os"

func atoi(s string) (int64, bool) {
	if s == "" {
		return 0, false
	}
	var n int64
	var neg bool
	start := 0
	if s[0] == '-' {
		neg = true
		start = 1
	} else if s[0] == '+' {
		start = 1
	}
	for i := start; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + int64(c-'0')
	}
	if neg {
		n = -n
	}
	return n, true
}

func putnbr(n int64) {
	if n < 0 {
		os.Stdout.Write([]byte("-"))
		n = -n
	}
	if n >= 10 {
		putnbr(n / 10)
	}
	os.Stdout.Write([]byte{byte('0' + n%10)})
}

func putstr(s string) {
	os.Stdout.Write([]byte(s))
}

// checkOverflowAdd returns false if a+b would overflow
func checkOverflowAdd(a, b int64) bool {
	if (b > 0 && a > 9223372036854775807-b) || (b < 0 && a < -9223372036854775808-b) {
		return false
	}
	return true
}

// checkOverflowSub returns false if a-b would overflow
func checkOverflowSub(a, b int64) bool {
	if (b < 0 && a > 9223372036854775807+b) || (b > 0 && a < -9223372036854775808+b) {
		return false
	}
	return true
}

// checkOverflowMul returns false if a*b would overflow
func checkOverflowMul(a, b int64) bool {
	if a == 0 || b == 0 {
		return true
	}
	if a == -1 && b == -9223372036854775808 {
		return false
	}
	if b == -1 && a == -9223372036854775808 {
		return false
	}
	if a > 9223372036854775807/b || a < -9223372036854775808/b {
		return false
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	a, ok1 := atoi(args[0])
	b, ok2 := atoi(args[2])
	op := args[1]

	if !ok1 || !ok2 {
		return
	}

	switch op {
	case "+":
		if !checkOverflowAdd(a, b) {
			return
		}
		putnbr(a + b)
	case "-":
		if !checkOverflowSub(a, b) {
			return
		}
		putnbr(a - b)
	case "*":
		if !checkOverflowMul(a, b) {
			return
		}
		putnbr(a * b)
	case "/":
		if b == 0 {
			putstr("No division by 0\n")
			return
		}
		putnbr(a / b)
	case "%":
		if b == 0 {
			putstr("No modulo by 0\n")
			return
		}
		putnbr(a % b)
	default:
		return
	}
	putstr("\n")
}
