package main

import (
	"fmt"
	"net"
	"os"
	"switchboard/internal"
)

func testHandler(conn net.Conn, server *internal.Server) {
	fmt.Println("Test Handler")
}

func main() {
	server, err := internal.CreateServer("localhost", 8080)
	if err != nil {
		fmt.Println("Could not start server", err)
		os.Exit(-1)
	}
	server.AcceptConnections(testHandler)
}
