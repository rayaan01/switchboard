package internal

import (
	"fmt"
	"io"
	"net"
	"strings"
)

var Store map[string]string = map[string]string{}

func Handler(conn net.Conn, s *Server) {
	for {
		buffer := make([]byte, 0, 1028)
		bytesRead := 0
		clientAddress := conn.RemoteAddr().String()
		err := read(&buffer, &bytesRead, conn)
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

func router(args []string) ([]byte, error) {
	errorMessage := []byte("Available commands:\n1. set [key] [value]\n2. get [key]\n3. exit \n")
	if len(args) == 0 {
		return errorMessage, nil
	}
	cmd := strings.ToLower(args[0])
	switch cmd {
	case "set":
		if len(args) != 3 {
			return errorMessage, nil
		}
		key := args[1]
		val := args[2]
		response, err := handleSet(key, val)
		if err != nil {
			return nil, err
		}
		return append(response, []byte("\n")...), nil
	case "get":
		if len(args) != 2 {
			return errorMessage, nil
		}
		key := args[1]
		response, err := handleGet(key)
		if err != nil {
			return nil, err
		}
		return append(response, []byte("\n")...), nil
	case "exit":
		return nil, io.EOF
	default:
		return errorMessage, nil
	}
}

func read(buffer *[]byte, bytesRead *int, conn net.Conn) error {
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

func handleSet(key string, value string) ([]byte, error) {
	Store[key] = value
	return []byte("OK"), nil
}

func handleGet(key string) ([]byte, error) {
	return []byte(Store[key]), nil
}
