package main

import (
	"bufio"
	"fmt"
	"os"
)

func welcomePrompt() {
	fmt.Println("Welcome to the switchboard CLI!")
	fmt.Println("You can use this tool to interact with the switchboard db. Usage -")
	fmt.Println("1. set [key] [value]")
	fmt.Println("2. get [key]")
	fmt.Println("3. del [key]")
	fmt.Println("4. exit")
}

func displayWrapper(text string) {
	fmt.Printf("%s %s", "switchboard >>", text)
}

func main() {
	welcomePrompt()
	displayWrapper("")
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		text := reader.Text()
		displayWrapper(text)
	}
}
