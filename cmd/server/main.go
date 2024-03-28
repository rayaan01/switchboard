package main

import (
	"fmt"
	"os"
	"switchboard/internal"
)

func main() {
	server, err := internal.CreateServer("localhost", 8080)
	if err != nil {
		fmt.Println("Could not start server", err)
		os.Exit(-1)
	}
	server.AcceptConnections(internal.Handler)
}
