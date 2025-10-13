package piscine

func IsPrimeA(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}

	i := 3
	for i*i <= nb {
		if nb%i == 0 {
			return false
		}
		i += 2
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	for {
		if IsPrime(nb) {
			return nb
		}
		nb++
	}
}
