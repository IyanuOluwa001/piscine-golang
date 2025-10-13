package piscine

func RecursiveFactorial(nb int) int {
    if nb < 0 {
        return 0
    }
    if nb == 0 || nb == 1 {
        return 1
    }

    prev := RecursiveFactorial(nb - 1)
    if prev == 0 {
        return 0 // propagate error (overflow or invalid)
    }

    result := nb * prev
    if result < 0 { // detect overflow
        return 0
    }
    return result
}
