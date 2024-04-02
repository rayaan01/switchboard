package main

import (
	"fmt"
	"os"

	"switchboard/pkg/db"
)

func main() {
	server, err := db.CreateServer("localhost", 8080)
	if err != nil {
		fmt.Printf("%s %s \n", "Could not start server", err)
		os.Exit(-1)
	}
	server.AcceptConnections(db.Handler)
}
