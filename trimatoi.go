package piscine

func TrimAtoi(s string) int {
	sign := 1
	numStarted := false
	num := 0
	signAllowed := true // sign allowed only before any digit

	for _, r := range s {
		if r == '-' && signAllowed && !numStarted {
			sign = -1
			signAllowed = false
		} else if r == '+' && signAllowed && !numStarted {
			// '+' sign found before digits, ignore sign, just disallow further signs
			signAllowed = false
		} else if r >= '0' && r <= '9' {
			numStarted = true
			signAllowed = false
			num = num*10 + int(r-'0')
		} else if numStarted && !(r >= '0' && r <= '9') {
			// once number started, if char is not digit, just ignore and continue
			continue
		} else {
			// non-digit, non-sign char before number, just continue scanning
			continue
		}
	}

	if !numStarted {
		return 0
	}
	return num * sign
}
