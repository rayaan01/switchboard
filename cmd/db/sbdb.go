package main

import (
	"flag"
	"fmt"
	"os"

	"switchboard/pkg/db"
)

func main() {
	host := flag.String("h", "localhost", "Host to serve on")
	port := flag.String("p", "8080", "Port to serve on")
	flag.Parse()
	server, err := db.CreateServer(*host, *port)
	if err != nil {
		fmt.Printf("%s %s \n", "Could not start server", err)
		os.Exit(-1)
	}
	server.AcceptConnections(db.Handler)
}
