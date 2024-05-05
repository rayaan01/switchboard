package db

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strings"
)

type Request struct {
	Key string
	Cmd string
}

func Handler(conn net.Conn, s *Server) {
	for {
		buffer := make([]byte, 0, 1024)
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
		var request Request
		json.Unmarshal(buffer[:bytesRead], &request)
		args := strings.Fields(request.Cmd)
		response, err := router(request.Key, args)
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
