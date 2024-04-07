package cli

import "fmt"

func WelcomePrompt() {
	fmt.Println("Welcome to the switchboard CLI!")
	fmt.Println("You can use this tool to interact with switchboard db. Usage -")
	fmt.Println("1. set [key] [value]")
	fmt.Println("2. get [key]")
	fmt.Println("3. del [key]")
	fmt.Println("4. exit")
	displayWrapper("")
}

func displayWrapper(text string) {
	fmt.Printf("%s %s", "switchboard >>", text)
}

func getUsageMessage() string {
	msg := "Available commands:\n1. set [key] [value]\n2. get [key]\n3. del [key]\n4. exit\n"
	return msg
}
