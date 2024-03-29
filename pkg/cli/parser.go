package cli

import (
	"fmt"
	"strings"
)

func Parser(input string) {
	args := strings.Fields(input)
	if len(args) == 0 {
		displayUsageMessage()
		return
	}
	cmd := args[0]
	switch cmd {
	case "set":
		fmt.Println("Set command is called")
	case "get":
		fmt.Println("Get command is called")
	case "del":
		fmt.Println("Del command is called")
	case "exit":
		fmt.Println("Exit command is called")
	}
}
