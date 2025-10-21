package main

import (
	"bufio"
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printError(err error) {
	printStr("ERROR: ")
	for _, r := range err.Error() {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			printStr(input)
		}
		return
	}

	for _, name := range args {
		file, err := os.Open(name)
		if err != nil {
			printError(err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, file)
		file.Close()

		if err != nil {
			printError(err)
			os.Exit(1)
		}
	}
}
