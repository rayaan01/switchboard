package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"switchboard/pkg/cli"
)

func main() {
	cli.WelcomePrompt()
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		input := reader.Text()
		response, err := cli.Parser(input)
		if err != nil {
			if err == io.EOF {
				fmt.Println("The connection has been closed. Thank you!")
				return
			}
			fmt.Println(err)
			continue
		}
		cli.DisplayWrapper(response)
	}
}
