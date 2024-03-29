package db

import (
	"fmt"
	"io"
	"net"
	"strings"
)

var Store = map[string]string{}

func Handler(conn net.Conn, s *Server) {
	for {
		buffer := make([]byte, 0, 1028)
		bytesRead := 0
		clientAddress := conn.RemoteAddr().String()
		err := readCommand(&buffer, &bytesRead, conn)
		if err != nil {
			if err == io.EOF {
				conn.Close()
				return
			}
			fmt.Printf("Could not read from connection %s : %s \n", clientAddress, err)
			conn.Write([]byte("Something went wrong!"))
			continue
		}
		input := string(buffer[:bytesRead-1])
		args := strings.Fields(input)
		response, err := router(args)
		if err != nil {
			if err == io.EOF {
				conn.Close()
				return
			}
			fmt.Printf("Router error on %s : %s \n", clientAddress, err)
			conn.Write([]byte("Something went wrong!"))
			continue
		}
		conn.Write(response)
	}
}

func readCommand(buffer *[]byte, bytesRead *int, conn net.Conn) error {
	for {
		chunk := make([]byte, 128)
		n, err := conn.Read(chunk)
		if err != nil {
			return err
		}
		*buffer = append(*buffer, chunk[:n]...)
		*bytesRead += n
		if n < cap(chunk) {
			return nil
		}
	}
}
