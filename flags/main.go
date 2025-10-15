package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	insertStr := ""
	orderFlag := false
	mainStr := ""

	// Parse arguments
	for _, arg := range args {
		if len(arg) >= 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
		} else if len(arg) >= 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			orderFlag = true
		} else {
			mainStr = arg
		}
	}

	result := mainStr + insertStr

	if orderFlag {
		result = sortString(result)
	}

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func sortString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[j] < runes[i] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}

func printHelp() {
	help := `--insert
  -i
         This flag inserts the string into the string passed as argument.
--order
  -o
         This flag will behave like a boolean, if it is called it will order the argument.`
	for _, r := range help {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
