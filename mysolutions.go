//Que 1: Write a function that returns the first rune of a string.
/*
main package

package piscine

func FirstRune1 (s string) rune {
	for _, r := range s{
		return r
	}
	return 0
}

q1p1
func FirstRune(s string) rune{
	for _, r:= range s{

	return r
	}

}

q1p2
func FirstRune1(s string) rune {
	for _, r := range s {
		z01.PrintRune(r) // Return the first rune found
	}
	z01.PrintRune('\n')

	for _, r:= range s{
		return r
	}
	return 0 // Return 0 if string is empty (no rune)
}


//Que 2: Write a function that returns the last rune of a string.

package piscine

func LastRune (s string) rune {
	var last rune
	for _, r := range s {
		last = r
	}
	return last
}

//Go gives us a shortcut for looping over collections like strings, arrays, slices, maps:
//for index, value := range something {
	// code
//}



//Que 3: Write a function that returns the nth rune of a string. If not possible, it returns 0.
package piscine

func NRune (s string, n int) rune {
	if n <= 0 {
		return 0
	}

	i := 1
	for _, r := range s {
		if i == n {
			return r
		}
		i++
	}

	return 0 // n is larger than the number of runes
}



//Que 4:
//Write a function that behaves like the Compare function.
package piscine

func Compare (a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}
*/

// Que 5:
// Write a function that counts the letters of a string and returns the count.
package piscine

func AlphaCount1(s string) int {
	iyanu := 0
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			iyanu++
		}
	}
	return iyanu
}

/*
Que 6:


func Index(s string, toFind string) int {
	if toFind == "" {
		return 0 // empty substring is always found at start
	}

	lenS := len(s)
	lenFind := len(toFind)

	if lenFind > lenS {
		return -1
	}

	for i := 0; i <= lenS-lenFind; i++ {
		if s[i:i+lenFind] == toFind {
			return i
		}
	}

	return -1
}

Que 7:

func Concat(str1 string, str2 string) string {
	return str1 + str2
}

Que 8:

func IsUpper(s string) bool {
	if s == "" {
		return false // empty string is not uppercase
	}
	for _, r := range s {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}

Que 9:

func IsLower(s string) bool {
	if s == "" {
		return false // empty string is not lowercase
	}
	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}

Que 10:

func IsAlpha(s string) bool {
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}

Que 11:

func IsNumeric(s string) bool {
	if s == "" {
		return false // empty string is not numeric
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

Que 12:

func IsPrintable(s string) bool {
	for _, r := range s {
		// Printable ASCII characters range from 32 (space) to 126 (~)
		if r < 32 || r > 126 {
			return false
		}
	}
	return true
}

Que 13:

func ToUpper(s string) string {
	result := []rune{}
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			result = append(result, r-('a'-'A'))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

Que 14:

func ToLower(s string) string {
	result := []rune{}
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			result = append(result, r+('a'-'A'))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

Que 15:


import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digitsCount := [10]int{}
	for n > 0 {
		digit := n % 10
		digitsCount[digit]++
		n /= 10
	}

	for digit := 0; digit <= 9; digit++ {
		for count := 0; count < digitsCount[digit]; count++ {
			z01.PrintRune(rune(digit + '0'))
		}
	}
}

Que 16:

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

Que 17:

func Capitalize(s string) string {
	result := []rune{}
	inWord := false

	for _, r := range s {
		if isAlnum(r) {
			if !inWord {
				// Start of a new word: capitalize if letter
				if r >= 'a' && r <= 'z' {
					r = r - ('a' - 'A')
				}
				inWord = true
			} else {
				// Inside a word: lowercase if uppercase letter
				if r >= 'A' && r <= 'Z' {
					r = r + ('a' - 'A')
				}
			}
		} else {
			inWord = false
		}
		result = append(result, r)
	}

	return string(result)
}

func isAlnum(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9')
}

Que 18:

func BasicJoin(elems []string) string {
	result := ""
	for _, str := range elems {
		result += str
	}
	return result
}

Que 19:

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}

	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result += sep + strs[i]
	}
	return result
}

Que 20:

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		// Handle corner case: nbr == MinInt (to avoid overflow)
		if nbr == -nbr {
			// MinInt is tricky in Go, but since int is at least 32-bit, we can use uint trick:
			// Cast to uint and handle as positive
			printBase(uint(-nbr), base)
			return
		}
		nbr = -nbr
	}

	printBase(uint(nbr), base)
}

func printBase(n uint, base string) {
	baseLen := uint(len(base))
	if n >= baseLen {
		printBase(n/baseLen, base)
	}
	z01.PrintRune(rune(base[n%baseLen]))
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' {
			return false
		}
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

Que 21:

func AtoiBase(s string, base string) int {
	// Check if base is valid
	if !isValidBase1(base) {
		return 0
	}

	baseLen := len(base)
	valueMap := make(map[rune]int)

	for i, char := range base {
		valueMap[char] = i
	}

	result := 0
	for _, char := range s {
		value, exists := valueMap[char]
		if !exists {
			return 0 // should never happen per instructions
		}
		result = result*baseLen + value
	}
	return result
}

// Helper to check base validity
func isValidBase1(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' {
			return false
		}
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

*/
