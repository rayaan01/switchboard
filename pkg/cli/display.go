package cli

import "fmt"

func WelcomePrompt() {
	fmt.Println("Welcome to the switchboard CLI!")
	fmt.Println("You can use this tool to interact with switchboard db. Usage -")
	fmt.Println("1. set [key] [value]")
	fmt.Println("2. get [key]")
	fmt.Println("3. del [key]")
	fmt.Println("4. get-range [lower_bound] [uppper_bound]")
	fmt.Println("5. create-access-key {HashTable|AVLTree}")
	fmt.Println("6. exit")
	displayWrapper("")
}

func displayWrapper(text string) {
	fmt.Printf("%s %s", "switchboard >>", text)
}
