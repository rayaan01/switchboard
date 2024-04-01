package main

import (
	"flag"
	"fmt"
	"os"

	"switchboard/pkg/cli"
)

func main() {
	cli.WelcomePrompt()
	host := flag.String("h", "localhost", "Host to connect to")
	port := flag.String("p", "8080", "Port to connect to")
	client, err := cli.CreateClient(host, port)
	if err != nil {
		fmt.Printf("%s %s \n", "Could not connect to server", err)
		os.Exit(-1)
	}
	client.HandleConnection()
}
