package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}

	n := start
	steps := 0
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		steps++
	}
	return steps
}
