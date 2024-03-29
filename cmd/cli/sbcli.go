package main

import (
	"bufio"
	"os"

	"switchboard/pkg/cli"
)

func main() {
	cli.WelcomePrompt()
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		input := reader.Text()
		cli.Parser(input)
		// displayWrapper(text)
	}
}
